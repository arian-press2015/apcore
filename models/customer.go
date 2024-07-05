package models

type Customer struct {
	BaseModel
	Name       string  `gorm:"not null" json:"name"`
	Details    string  `json:"details"`
	Phone      string  `gorm:"unique;not null" json:"phone"`
	Logo       string  `json:"logo"`
	IsActive   bool    `json:"isActive"`
	IsDisabled bool    `json:"isDisabled"`
	Latitude   float64 `gorm:"not null" json:"latitude"`
	Longitude  float64 `gorm:"not null" json:"longitude"`
}
