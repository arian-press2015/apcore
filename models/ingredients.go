package models

type Ingredient struct {
	BaseModel
	Name         string `gorm:"not null" json:"name"`
	ImageAddress string `gorm:"not null" json:"image_address"`
}
