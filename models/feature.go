package models

type Feature struct {
	BaseModel
	Name        string `gorm:"unique;not null" json:"name" binding:"required"`
	Description string `gorm:"not null" json:"description" binding:"required"`
	Enabled     bool   `gorm:"not null" json:"enabled" binding:"required"`
}
