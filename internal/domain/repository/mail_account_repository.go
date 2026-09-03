package repository

import "github.com/daexaf/mailvault/internal/domain/entities"

type MailAccountRepository interface {
	Create(account *entities.MailAccount) error
	FindAll() ([]entities.MailAccount, error)
	FindByID(id uint) (*entities.MailAccount, error)
	ExistsByEmail(email string) (bool, error)
	UpdateLastUID(id uint, lastUID uint32) error
}
