package service

import (
	"github.com/daexaf/mailvault/internal/domain/repository"
	"github.com/daexaf/mailvault/internal/dto"
)

type EmailService struct {
	repo repository.EmailRepository
}

func NewEmailService(
	repo repository.EmailRepository,
) *EmailService {
	return &EmailService{
		repo: repo,
	}
}

func (s *EmailService) Search(
	subject string,
	from string,
	to string,
	keyword string,
) ([]dto.EmailResponse, error) {

	emails, err := s.repo.Search(
		subject,
		from,
		to,
		keyword,
	)

	if err != nil {
		return nil, err
	}

	result := make(
		[]dto.EmailResponse,
		0,
		len(emails),
	)

	for _, email := range emails {
		result = append(
			result,
			dto.EmailResponse{
				ID:            email.ID,
				MailAccountID: email.MailAccountID,
				MessageID:     email.ProviderMessageID,
				Subject:       email.Subject,
				From:          email.From,
				To:            email.To,
				CC:            email.CC,
				BCC:           email.BCC,
				ReceivedAt:    email.ReceivedAt,
				HasAttachment: email.HasAttachment,
				Folder:        email.Folder,
			},
		)
	}

	return result, nil
}
