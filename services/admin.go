package services

import (
	"apcore/models"
	"apcore/repositories"

	"github.com/google/uuid"
)

type AdminService interface {
	CreateAdmin(admin *models.Admin) error
	GetAdmins(offset int, limit int) ([]models.Admin, error)
	GetAdminByID(id uuid.UUID) (*models.Admin, error)
	GetAdminByName(name string) (*models.Admin, error)
	GetAdminByPhone(phone string) (*models.Admin, error)
	UpdateAdmin(admin *models.Admin) error
	DeleteAdmin(id uuid.UUID) error
}

type adminService struct {
	repo repositories.AdminRepository
}

func NewAdminService(repo repositories.AdminRepository) AdminService {
	return &adminService{repo}
}

func (s *adminService) CreateAdmin(admin *models.Admin) error {
	return s.repo.CreateAdmin(admin)
}

func (s *adminService) GetAdmins(offset int, limit int) ([]models.Admin, error) {
	return s.repo.GetAdmins(offset, limit)
}

func (s *adminService) GetAdminByID(id uuid.UUID) (*models.Admin, error) {
	return s.repo.GetAdminByID(id)
}

func (s *adminService) GetAdminByName(name string) (*models.Admin, error) {
	return s.repo.GetAdminByName(name)
}

func (s *adminService) GetAdminByPhone(phone string) (*models.Admin, error) {
	return s.repo.GetAdminByPhone(phone)
}

func (s *adminService) UpdateAdmin(admin *models.Admin) error {
	return s.repo.UpdateAdmin(admin)
}

func (s *adminService) DeleteAdmin(id uuid.UUID) error {
	return s.repo.DeleteAdmin(id)
}
