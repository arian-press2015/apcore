package repositories

import (
	"apcore/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetCategories(offset int, limit int) ([]models.Category, error)
	GetCategoryCount() (int64, error)
	GetCategoryByID(uuid string) (*models.Category, error)
	GetCategoryBySlug(slug string, customerID uuid.UUID) (*models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uuid.UUID) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetCategories(offset int, limit int) ([]models.Category, error) {
	var categorys []models.Category
	err := r.db.Offset(offset).Limit(limit).Find(&categorys).Error
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func (r *categoryRepository) GetCategoryCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.Category{}).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *categoryRepository) GetCategoryByID(uuid string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ?", uuid).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetCategoryBySlug(slug string, customerID uuid.UUID) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("slug = ? AND customer_id = ?", slug, customerID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) UpdateCategory(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(id uuid.UUID) error {
	return r.db.Delete(&models.Category{}, id).Error
}
