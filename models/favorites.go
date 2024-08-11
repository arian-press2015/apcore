package models

import "github.com/google/uuid"

type Favorites struct {
	BaseModel
	CustomerID uuid.UUID `gorm:"type:uuid;not null" json:"customer_id"`
	Customer   Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
}
