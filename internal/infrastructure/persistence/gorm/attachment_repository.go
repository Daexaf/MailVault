package gorm

import (
	"github.com/daexaf/mailvault/internal/domain/entities"
	"github.com/daexaf/mailvault/internal/domain/repository"
	"gorm.io/gorm"
)

type attachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) repository.AttachmentRepository {
	return &attachmentRepository{
		db: db,
	}
}

func (r *attachmentRepository) Create(
	attachment *entities.Attachment,
) error {
	return r.db.Create(attachment).Error
}
