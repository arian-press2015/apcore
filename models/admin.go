package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
}
