package models

import "github.com/google/uuid"

type Menu struct {
	BaseModel
	CustomerID  uuid.UUID `gorm:"type:uuid;not null" json:"customer_id"`
	Categories []Category `gorm:"foreignKey:MenuID" json:"categories"`
}
