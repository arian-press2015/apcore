package models

import "github.com/google/uuid"

type Role struct {
	BaseModel
	Name        string       `gorm:"unique;not null" json:"name" binding:"required"`
	Users       []User       `gorm:"many2many:user_roles;" json:"users"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type UserRole struct {
    UserID uuid.UUID
    RoleID uuid.UUID
}

type RolePermission struct {
    RoleID       uuid.UUID
    PermissionID uuid.UUID
}
