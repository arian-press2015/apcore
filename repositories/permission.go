package repositories

import (
	"apcore/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	CreatePermission(permission *models.Permission) error
	GetPermissions(offset int, limit int) ([]models.Permission, error)
	GetPermissionByID(id uuid.UUID) (*models.Permission, error)
	GetPermissionByName(Name string) (*models.Permission, error)
	UpdatePermission(permission *models.Permission) error
	DeletePermission(id uuid.UUID) error
	HasUserSomePermission(userID uuid.UUID, permissionName string) (bool, error)
}
type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db}
}

func (r *permissionRepository) CreatePermission(permission *models.Permission) error {
	return r.db.Create(permission).Error
}

func (r *permissionRepository) GetPermissions(offset int, limit int) ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.Offset(offset).Limit(limit).Preload("Users").Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *permissionRepository) GetPermissionByID(id uuid.UUID) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.First(&permission, id).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) GetPermissionByName(name string) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.Where("name = ?", name).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) UpdatePermission(permission *models.Permission) error {
	return r.db.Save(permission).Error
}

func (r *permissionRepository) DeletePermission(id uuid.UUID) error {
	return r.db.Delete(&models.Permission{}, id).Error
}

func (r *permissionRepository) HasUserSomePermission(userID uuid.UUID, permissionName string) (bool, error) {
	var count int64

	err := r.db.Table("roles").
		Select("COUNT(roles.id)").
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Joins("JOIN role_permissions ON role_permissions.role_id = roles.id").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Where("user_roles.user_id = ? AND permissions.name = ?", userID, permissionName).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
