package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
}
