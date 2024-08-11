package models

import "github.com/google/uuid"

type Customer struct {
	BaseModel
	Name       string  `gorm:"not null" json:"name"`
	Slug       string  `gorm:"unique;not null" json:"slug"`
	Details    string  `json:"details"`
	Phone      string  `gorm:"unique;not null" json:"phone"`
	Logo       string  `json:"logo"`
	IsActive   bool    `json:"isActive"`
	IsDisabled bool    `json:"isDisabled"`
	Latitude   float64 `gorm:"not null" json:"latitude"`
	Longitude  float64 `gorm:"not null" json:"longitude"`
	Users      []User  `gorm:"many2many:user_customers;" json:"users"`
}

type UserCustomer struct {
	UserID     uuid.UUID
	CustomerID uuid.UUID
}
