package repositories

import (
	"apcore/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRepository interface {
	GetMenu(customerID uuid.UUID) (*models.Menu, error)
	CreateMenu(menu *models.Menu) error
	UpdateMenu(menu *models.Menu) error
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db}
}

func (r *menuRepository) GetMenu(customerID uuid.UUID) (*models.Menu, error) {
	var menu models.Menu

	err := r.db.Where("customer_id = ?", customerID).First(&menu).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) CreateMenu(menu *models.Menu) error {
	return r.db.Create(menu).Error
}

func (r *menuRepository) UpdateMenu(menu *models.Menu) error {
	return r.db.Save(menu).Error
}
