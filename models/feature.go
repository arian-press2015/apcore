package models

import (
	"gorm.io/gorm"
)

type Feature struct {
	gorm.Model
	Name        string `gorm:"unique;not null" json:"name" binding:"required"`
	Description string `gorm:"not null" json:"description" binding:"required"`
	Enabled     bool   `gorm:"not null" json:"enabled" binding:"required"`
}
