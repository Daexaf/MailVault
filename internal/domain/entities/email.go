package entities

import (
	"time"
)

type Email struct {
	BaseEntity

	MailAccountID     uint        `gorm:"not null;index"`
	MailAccount       MailAccount `gorm:"foreignKey:MailAccountID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
	ProviderMessageID string      `gorm:"size:255;not null;uniqueIndex"`
	To                string      `gorm:"size:255;not null"`
	From              string      `gorm:"size:255;not null"`
	Subject           string      `gorm:"size:255;not null"`
	Body              string      `gorm:"type:text;not null"`
	CC                string      `gorm:"size:255"`
	BCC               string      `gorm:"size:255"`
	ReceivedAt        time.Time   `gorm:"not null;index"`
	HasAttachment     bool        `gorm:"not null; default:false"`
	Folder            string      `gorm:"size:20;not null;index"`
}
