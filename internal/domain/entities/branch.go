package entities

type Branch struct {
	BaseEntity

	Code     string `gorm:size:20;not null;uniqueIndex"`
	Name     string `gorm:size:100;not null"`
	IsActive bool   `gorm:"default:true"`
}
