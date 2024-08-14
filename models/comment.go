package models

import "github.com/google/uuid"

type Comment struct {
	BaseModel
	Title     string    `gorm:"not null" json:"title"`
	Body      string    `gorm:"not null" json:"body"`
	Rate      int       `gorm:"not null" json:"rate"`
	ProductID uuid.UUID `gorm:"type:uuid" json:"product_id"`
}
