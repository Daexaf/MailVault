package entities

import (
	"time"
)

type MailAccount struct {
	BaseEntity

	BranchID uint   `gorm:"not null; index"`
	Branch   Branch `gorm:"foreignKey:BranchID; constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`

	Email    string `gorm:"size:255;not null; uniqueIndex"`
	Password string `gorm:"size:255;not null"`

	Host   string `gorm:"size:255;not null"`
	Port   int    `gorm:"not null"`
	UseSSL bool   `gorm:"not null; default:true"`

	IsActive   bool `gorm:"not null; default:true"`
	LastSyncAt *time.Time
}
