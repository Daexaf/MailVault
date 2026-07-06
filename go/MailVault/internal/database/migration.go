package database

import (
	"fmt"

	"github.com/daexaf/mailvault/internal/domain/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.Branch{},
		&entities.MailAccount{},
	)

	if err != nil {
		return fmt.Errorf("failed to migrate  auto: %w", err)
	}

	return nil
}
