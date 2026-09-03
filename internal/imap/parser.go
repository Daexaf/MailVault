package imap

import (
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"net/mail"
	"strings"
)

type ParsedMessage struct {
	Body          string
	HasAttachment bool
	Attachment    []ParsedAttachment
}

type ParsedAttachment struct {
	FileName    string
	ContentType string
	Data        []byte
}

func ParseRawMessage(raw string) (*ParsedMessage, error) {
	message, err := mail.ReadMessage(strings.NewReader(raw))
	if err != nil {
		return nil, err
	}

	contentType := message.Header.Get("Content-Type")

	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		// Beberapa email lama/malformed tidak punya Content-Type yang rapi.
		bodyBytes, readErr := io.ReadAll(message.Body)
		if readErr != nil {
			return nil, readErr
		}

		return &ParsedMessage{
			Body: string(bodyBytes),
		}, nil
	}

	result := &ParsedMessage{}

	// Email multipart
	if strings.HasPrefix(mediaType, "multipart/") {
		boundary := params["boundary"]

		reader := multipart.NewReader(message.Body, boundary)

		if err := parseMultipart(reader, result); err != nil {
			return nil, err
		}

		return result, nil
	}

	// Email singlepart
	bodyBytes, err := readDecodedBody(
		message.Body,
		message.Header.Get("Content-Transfer-Encoding"),
	)
	if err != nil {
		return nil, err
	}

	result.Body = string(bodyBytes)

	return result, nil
}

func parseMultipart(
	reader *multipart.Reader,
	result *ParsedMessage,
) error {
	for {
		part, err := reader.NextPart()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		contentType := part.Header.Get("Content-Type")
		disposition := part.Header.Get("Content-Disposition")
		transferEncoding := part.Header.Get("Content-Transfer-Encoding")

		mediaType, params, err := mime.ParseMediaType(contentType)
		if err != nil {
			mediaType = contentType
			params = map[string]string{}
		}

		// Cek apakah part ini attachment
		dispositionType, dispositionParams, _ :=
			mime.ParseMediaType(disposition)

		filename := dispositionParams["filename"]

		// Beberapa email menyimpan nama file di Content-Type:
		// image/jpeg; name="image.jpg"
		if filename == "" {
			filename = params["name"]
		}

		if dispositionType == "attachment" || filename != "" {
			result.HasAttachment = true

			fileBytes, err := readDecodedBody(
				part,
				transferEncoding,
			)
			if err != nil {
				return err
			}

			result.Attachment = append(
				result.Attachment,
				ParsedAttachment{
					FileName:    filename,
					ContentType: mediaType,
					Data:        fileBytes,
				},
			)

			continue
		}

		// Multipart di dalam multipart
		if strings.HasPrefix(mediaType, "multipart/") {
			boundary := params["boundary"]

			if boundary == "" {
				continue
			}

			nestedReader := multipart.NewReader(
				part,
				boundary,
			)

			if err := parseMultipart(
				nestedReader,
				result,
			); err != nil {
				return err
			}

			continue
		}

		bodyBytes, err := readDecodedBody(
			part,
			transferEncoding,
		)
		if err != nil {
			return err
		}

		// Prioritaskan text/plain daripada text/html
		if mediaType == "text/plain" {
			if result.Body == "" {
				result.Body = string(bodyBytes)
			}

			continue
		}

		// Kalau tidak ada text/plain, ambil text/html
		if mediaType == "text/html" && result.Body == "" {
			result.Body = string(bodyBytes)
		}
	}

	return nil
}

func readDecodedBody(
	reader io.Reader,
	transferEncoding string,
) ([]byte, error) {
	switch strings.ToLower(
		strings.TrimSpace(transferEncoding),
	) {

	case "base64":
		return io.ReadAll(
			base64.NewDecoder(
				base64.StdEncoding,
				reader,
			),
		)

	case "quoted-printable":
		return io.ReadAll(
			quotedprintable.NewReader(reader),
		)

	default:
		return io.ReadAll(reader)
	}
}
