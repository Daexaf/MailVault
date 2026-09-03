package dto

import "time"

type EmailResponse struct {
	ID            uint      `json:"id"`
	MailAccountID uint      `json:"mail_account_id"`
	MessageID     string    `json:"message_id"`
	Subject       string    `json:"subject"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	CC            string    `json:"cc"`
	BCC           string    `json:"bcc"`
	ReceivedAt    time.Time `json:"received_at"`
	HasAttachment bool      `json:"has_attachment"`
	Folder        string    `json:"folder"`
}
