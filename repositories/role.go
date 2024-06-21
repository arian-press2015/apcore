package repositories

import (
	"apcore/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role *models.Role) error
	GetRoles(offset int, limit int) ([]models.Role, error)
	GetRoleByID(id uint) (*models.Role, error)
	GetRoleByName(Name string) (*models.Role, error)
	UpdateRole(role *models.Role) error
	DeleteRole(id uint) error
}
type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) CreateRole(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetRoles(offset int, limit int) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Offset(offset).Limit(limit).Preload("Users").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) GetRoleByID(id uint) (*models.Role, error) {
	var role models.Role
	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) UpdateRole(role *models.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) DeleteRole(id uint) error {
	return r.db.Delete(&models.Role{}, id).Error
}
