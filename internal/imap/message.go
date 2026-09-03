package imap

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

var (
	ErrInboxEmpty   = errors.New("inbox is empty")
	ErrMessageEmpty = errors.New("message is empty")
)

type Message struct {
	MessageID     string
	Subject       string
	From          string
	To            string
	CC            string
	BCC           string
	Date          time.Time
	Body          string
	HasAttachment bool
	Attachments   []ParsedAttachment
}

type MessageMetadata struct {
	UID       uint32
	MessageID string
	Subject   string
	From      string
	To        string
	Date      time.Time
}

func FetchMessageMetadataByUID(
	host string,
	port int,
	email string,
	password string,
	lastUID uint32,
	batchSize uint32,
) ([]MessageMetadata, error) {
	address := fmt.Sprintf("%s:%d", host, port)

	c, err := client.DialTLS(address, &tls.Config{})
	if err != nil {
		return nil, err
	}
	defer c.Logout()

	if err := c.Login(email, password); err != nil {
		return nil, err
	}

	mbox, err := c.Select("INBOX", true)
	if err != nil {
		return nil, err
	}

	if mbox.Messages == 0 {
		return []MessageMetadata{}, nil
	}

	//cari uid setelah uid terakhir yang sudah berhasil disimpan
	criteria := imap.NewSearchCriteria()

	if lastUID > 0 {
		criteria.Uid = new(imap.SeqSet)
		criteria.Uid.AddRange(
			lastUID+1,
			0,
		)
	}

	uids, err := c.UidSearch(criteria)
	if err != nil {
		return nil, err
	}

	if len(uids) == 0 {
		return []MessageMetadata{}, nil
	}

	if uint32(len(uids)) > batchSize {
		uids = uids[:batchSize]
	}

	// seqSet := new(imap.SeqSet)
	// seqSet.AddRange(1, mbox.Messages)
	uidSet := new(imap.SeqSet)

	for _, uid := range uids {
		uidSet.AddNum(uid)
	}

	items := []imap.FetchItem{
		imap.FetchEnvelope,
		imap.FetchUid,
	}

	messages := make(chan *imap.Message, len(uids))
	done := make(chan error, 1)

	go func() {
		done <- c.Fetch(
			uidSet,
			items,
			messages,
		)
	}()

	result := make([]MessageMetadata, 0)

	for msg := range messages {
		if msg == nil || msg.Envelope == nil {
			continue
		}

		result = append(
			result,
			MessageMetadata{
				UID:       msg.Uid,
				MessageID: msg.Envelope.MessageId,
				Subject:   msg.Envelope.Subject,
				From:      addressesToString(msg.Envelope.From),
				To:        addressesToString(msg.Envelope.To),
				Date:      msg.Envelope.Date,
			},
		)
	}

	if err := <-done; err != nil {
		return nil, err
	}

	return result, nil
}

func FetchMessageByUID(
	host string,
	port int,
	email string,
	password string,
	uid uint32,
) (*Message, error) {

	address := fmt.Sprintf("%s:%d", host, port)

	// Connect ke IMAP server
	c, err := client.DialTLS(
		address,
		&tls.Config{},
	)
	if err != nil {
		return nil, err
	}

	defer c.Logout()

	// Login
	if err := c.Login(email, password); err != nil {
		return nil, err
	}

	// Buka INBOX read-only
	_, err = c.Select("INBOX", true)
	if err != nil {
		return nil, err
	}

	// Ambil hanya email dengan sequence number
	// yang diberikan dari metadata.
	uidSet := new(imap.SeqSet)
	uidSet.AddNum(uid)

	// Kita membutuhkan body email
	// agar bisa diparse body + attachment.
	section := &imap.BodySectionName{
		Peek: true,
	}

	items := []imap.FetchItem{
		imap.FetchEnvelope,
		imap.FetchUid,
		section.FetchItem(),
	}

	messages := make(chan *imap.Message, 1)
	done := make(chan error, 1)

	go func() {
		done <- c.UidFetch(
			uidSet,
			items,
			messages,
		)
	}()

	var result *Message

	for msg := range messages {

		if msg == nil || msg.Envelope == nil {
			continue
		}

		// =============================
		// RAW BODY
		// =============================

		rawBody := ""

		bodyReader := msg.GetBody(section)

		if bodyReader != nil {

			bodyBytes, err := io.ReadAll(
				bodyReader,
			)

			if err != nil {
				return nil, err
			}

			rawBody = string(bodyBytes)
		}

		// =============================
		// PARSE BODY + ATTACHMENT
		// =============================

		body := ""
		hasAttachment := false

		var attachments []ParsedAttachment

		if rawBody != "" {

			parsed, err := ParseRawMessage(
				rawBody,
			)

			if err != nil {
				return nil, err
			}

			body = parsed.Body

			hasAttachment =
				parsed.HasAttachment

			attachments =
				parsed.Attachment
		}

		// =============================
		// CREATE RESULT
		// =============================

		result = &Message{
			MessageID: msg.Envelope.MessageId,

			Subject: msg.Envelope.Subject,

			From: addressesToString(
				msg.Envelope.From,
			),

			To: addressesToString(
				msg.Envelope.To,
			),

			CC: addressesToString(
				msg.Envelope.Cc,
			),

			BCC: addressesToString(
				msg.Envelope.Bcc,
			),

			Date: msg.Envelope.Date,

			Body: body,

			HasAttachment: hasAttachment,

			Attachments: attachments,
		}
	}

	if err := <-done; err != nil {
		return nil, err
	}

	if result == nil {
		return nil, ErrMessageEmpty
	}

	return result, nil
}

func addressesToString(addresses []*imap.Address) string {
	var result []string

	for _, addr := range addresses {
		if addr == nil {
			continue
		}

		if addr.MailboxName == "" || addr.HostName == "" {
			continue
		}

		email := fmt.Sprintf("%s@%s", addr.MailboxName, addr.HostName)

		result = append(result, email)
	}

	return strings.Join(result, ", ")
}
