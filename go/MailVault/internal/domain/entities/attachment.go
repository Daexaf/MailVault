package entities

type Attachment struct {
	BaseEntity

	EmailID      uint   `gorm:"not null;index"`
	Email        Email  `gorm:"foreignKey:EmailID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OriginalName string `gorm:"size:255;not null"`
	StoredName   string `gorm:"size:255;not null"`
	StoragePath  string `gorm:"size:255;not null"`
	FileSize     int64  `gorm:"not null"`
	MimeType     string `gorm:"size:100;not null"`
}
