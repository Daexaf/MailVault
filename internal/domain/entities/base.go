package entities

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
