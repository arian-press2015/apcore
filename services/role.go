package services

import (
	"apcore/models"
	"apcore/repositories"

	"github.com/google/uuid"
)

type RoleService interface {
	CreateRole(role *models.Role) error
	GetRoles(offset int, limit int) ([]models.Role, error)
	GetRoleByID(id uuid.UUID) (*models.Role, error)
	GetRoleByName(Name string) (*models.Role, error)
	UpdateRole(role *models.Role) error
	DeleteRole(id uuid.UUID) error
}

type roleService struct {
	repo repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) RoleService {
	return &roleService{repo}
}

func (s *roleService) CreateRole(role *models.Role) error {
	return s.repo.CreateRole(role)
}

func (s *roleService) GetRoles(offset int, limit int) ([]models.Role, error) {
	return s.repo.GetRoles(offset, limit)
}

func (s *roleService) GetRoleByID(id uuid.UUID) (*models.Role, error) {
	return s.repo.GetRoleByID(id)
}

func (s *roleService) GetRoleByName(name string) (*models.Role, error) {
	return s.repo.GetRoleByName(name)
}

func (s *roleService) UpdateRole(role *models.Role) error {
	return s.repo.UpdateRole(role)
}

func (s *roleService) DeleteRole(id uuid.UUID) error {
	return s.repo.DeleteRole(id)
}
