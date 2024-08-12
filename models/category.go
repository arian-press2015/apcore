package models

import "github.com/google/uuid"

type Category struct {
	BaseModel
	Name               string     `gorm:"not null" json:"name"`
	Slug               string     `gorm:"unique;not null" json:"slug"`
	MenuID             uuid.UUID  `gorm:"type:uuid;not null" json:"menu_id"`
	CategoryTemplateID *uuid.UUID `gorm:"type:uuid" json:"category_template_id"`
	Products           []Product  `gorm:"foreignKey:CategoryID" json:"products"`
}
