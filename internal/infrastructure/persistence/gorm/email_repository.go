package gorm

import (
	"github.com/daexaf/mailvault/internal/domain/entities"
	"github.com/daexaf/mailvault/internal/domain/repository"
	"gorm.io/gorm"
)

type emailRepository struct {
	db *gorm.DB
}

func NewEmailRepository(db *gorm.DB) repository.EmailRepository {
	return &emailRepository{
		db: db,
	}
}

func (r *emailRepository) Create(email *entities.Email) error {
	return r.db.Create(email).Error
}

func (r *emailRepository) ExistByProviderMessageID(providerMessageID string) (bool, error) {
	var count int64

	err := r.db.Model(&entities.Email{}).
		Where("provider_message_id = ?", providerMessageID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *emailRepository) Delete(email *entities.Email) error {
	return r.db.Delete(email).Error
}

func (r *emailRepository) DeletePermanently(email *entities.Email) error {
	return r.db.Unscoped().Delete(email).Error
}

func (r *emailRepository) Search(
	subject string,
	from string,
	to string,
	keyword string,
) ([]entities.Email, error) {

	var emails []entities.Email

	query := r.db.Model(&entities.Email{})

	if subject != "" {
		query = query.Where(
			"subject LIKE ?",
			"%"+subject+"%",
		)
	}

	if from != "" {
		query = query.Where(
			"[from] LIKE ?",
			"%"+from+"%",
		)
	}

	if to != "" {
		query = query.Where(
			"[to] LIKE ?",
			"%"+to+"%",
		)
	}

	if keyword != "" {
		search := "%" + keyword + "%"

		query = query.Where(
			"(subject LIKE ? OR body LIKE ? OR [from] LIKE ? OR [to] LIKE ?)",
			search,
			search,
			search,
			search,
		)
	}

	err := query.
		Order("received_at DESC").
		Find(&emails).
		Error

	if err != nil {
		return nil, err
	}

	return emails, nil
}
