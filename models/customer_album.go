package models

import "github.com/google/uuid"

type CustomerAlbum struct {
	BaseModel
	Name    string    `gorm:"not null" json:"name"`
	OwnerId uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`
	Owner   Customer  `gorm:"foreignKey:OwnerId" json:"owner"`
	Address string    `gorm:"not null" json:"address"`
}
