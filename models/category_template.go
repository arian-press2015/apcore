package models

type CategoryTemplate struct {
	BaseModel
	Name        string            `gorm:"not null" json:"name"`
	Slug        string            `gorm:"unique;not null" json:"slug"`
	Description string            `gorm:"not null" json:"description"`
	Products    []TemplateProduct `gorm:"foreignKey:TemplateID" json:"products"`
}
