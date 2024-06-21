package models

type Role struct {
	BaseModel
	Name  string `gorm:"unique;not null" json:"name" binding:"required"`
	Users []User `gorm:"many2many:user_roles;"`
}
