package repository

import "github.com/daexaf/mailvault/internal/domain/entities"

type EmailRepository interface {
	Create(email *entities.Email) error
	ExistByProviderMessageID(providerMessageID string) (bool, error)
	DeletePermanently(email *entities.Email) error
	Delete(email *entities.Email) error

	Search(
		subject string,
		from string,
		to string,
		keyword string,
	) ([]entities.Email, error)
}
