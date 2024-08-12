package models

import "github.com/google/uuid"

type Product struct {
	BaseModel
	Name              string       `gorm:"not null" json:"name"`
	Slug              string       `gorm:"unique;not null" json:"slug"`
	Price             float64      `gorm:"not null" json:"price"`
	CategoryID        uuid.UUID    `gorm:"type:uuid;not null" json:"category_id"`
	ProductTemplateID *uuid.UUID   `gorm:"type:uuid" json:"product_template_id"`
	Ingredients       []Ingredient `gorm:"many2many:product_ingredients;" json:"ingredients"`
}

type ProductIngredient struct {
	ProductID    uuid.UUID `gorm:"type:uuid;not null"`
	IngredientID uuid.UUID `gorm:"type:uuid;not null"`
}
