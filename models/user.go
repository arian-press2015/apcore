package models

type User struct {
	BaseModel
	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Roles    []Role `gorm:"many2many:user_roles;" json:"roles" binding:"required"`
}
