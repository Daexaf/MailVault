package repository

import "github.com/daexaf/mailvault/internal/domain/entities"

type AttachmentRepository interface {
	Create(attachment *entities.Attachment) error
}
