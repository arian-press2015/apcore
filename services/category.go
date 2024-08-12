package services

import (
	"apcore/models"
	"apcore/repositories"

	"github.com/google/uuid"
)

type CategoryService interface {
	CreateCategory(category *models.Category) error
	GetCategories(offset int, limit int) ([]models.Category, error)
	GetCategoryCount() (int64, error)
	GetCategoryByID(uuid string) (*models.Category, error)
	GetCategoryBySlug(slug string, customerID uuid.UUID) (*models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uuid.UUID) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	return s.repo.CreateCategory(category)
}

func (s *categoryService) GetCategories(offset int, limit int) ([]models.Category, error) {
	return s.repo.GetCategories(offset, limit)
}

func (s *categoryService) GetCategoryCount() (int64, error) {
	return s.repo.GetCategoryCount()
}

func (s *categoryService) GetCategoryByID(uuid string) (*models.Category, error) {
	return s.repo.GetCategoryByID(uuid)
}

func (s *categoryService) GetCategoryBySlug(slug string, customerID uuid.UUID) (*models.Category, error) {
	return s.repo.GetCategoryBySlug(slug, customerID)
}

func (s *categoryService) UpdateCategory(category *models.Category) error {
	return s.repo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(id uuid.UUID) error {
	return s.repo.DeleteCategory(id)
}
