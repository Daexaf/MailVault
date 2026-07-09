package gorm

import (
	"github.com/daexaf/mailvault/internal/domain/entities"
	"github.com/daexaf/mailvault/internal/domain/repository"
	"gorm.io/gorm"
)

type mailAccountRepository struct {
	db *gorm.DB
}

func NewMailAccountRepository(db *gorm.DB) repository.MailAccountRepository {
	return &mailAccountRepository{
		db: db,
	}
}

func (r *mailAccountRepository) Create(account *entities.MailAccount) error {
	return r.db.Create(account).Error
}

func (r *mailAccountRepository) ExistsByEmail(email string) (bool, error) {
	var count int64

	err := r.db.Model(&entities.MailAccount{}).
		Where("email = ?", email).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *mailAccountRepository) FindAll() ([]entities.MailAccount, error) {
	var accounts []entities.MailAccount

	if err := r.db.Find(&accounts).Error; err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *mailAccountRepository) FindByID(id uint) (*entities.MailAccount, error) {
	var account entities.MailAccount

	if err := r.db.First(&account, id).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
