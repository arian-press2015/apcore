package models

type CategoryTemplate struct {
	BaseModel
	Name        string            `gorm:"not null" json:"name"`
	Description string            `gorm:"not null" json:"description"`
	Products    []TemplateProduct `gorm:"foreignKey:TemplateID" json:"products"`
}
