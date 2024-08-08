package repositories

import (
	"apcore/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(admin *models.Admin) error
	GetAdmins(offset int, limit int) ([]models.Admin, error)
	GetAdminByID(id uuid.UUID) (*models.Admin, error)
	GetAdminByName(name string) (*models.Admin, error)
	GetAdminByPhone(phone string) (*models.Admin, error)
	UpdateAdmin(admin *models.Admin) error
	DeleteAdmin(id uuid.UUID) error
}
type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) CreateAdmin(admin *models.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) GetAdmins(offset int, limit int) ([]models.Admin, error) {
	var admins []models.Admin
	err := r.db.Offset(offset).Limit(limit).Preload("Users").Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *adminRepository) GetAdminByID(id uuid.UUID) (*models.Admin, error) {
	var admin models.Admin
	err := r.db.First(&admin, id).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) GetAdminByName(name string) (*models.Admin, error) {
	var admin models.Admin
	err := r.db.Where("username = ?", name).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) GetAdminByPhone(phone string) (*models.Admin, error) {
	var admin models.Admin
	err := r.db.Where("phone = ?", phone).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) UpdateAdmin(admin *models.Admin) error {
	return r.db.Save(admin).Error
}

func (r *adminRepository) DeleteAdmin(id uuid.UUID) error {
	return r.db.Delete(&models.Admin{}, id).Error
}
