package models

import "github.com/google/uuid"

type TemplateProduct struct {
	BaseModel
	Name        string       `gorm:"not null" json:"name"`
	Description string       `gorm:"not null" json:"description"`
	TemplateID  uuid.UUID    `gorm:"type:uuid;not null" json:"template_id"`
	Ingredients []Ingredient `gorm:"many2many:template_product_ingredients;" json:"ingredients"`
}

type TemplateProductIngredient struct {
	TemplateProductID uuid.UUID `gorm:"type:uuid;not null"`
	IngredientID      uuid.UUID `gorm:"type:uuid;not null"`
}