package services

import (
	"apcore/models"
	"apcore/repositories"

	"github.com/google/uuid"
)

type IngredientService interface {
	CreateIngredient(ingredient *models.Ingredient) error
	GetIngredients(offset int, limit int) ([]models.Ingredient, error)
	GetIngredientCount() (int64, error)
	GetIngredientByID(uuid string) (*models.Ingredient, error)
	UpdateIngredient(ingredient *models.Ingredient) error
	DeleteIngredient(id uuid.UUID) error
}

type ingredientService struct {
	repo repositories.IngredientRepository
}

func NewIngredientService(repo repositories.IngredientRepository) IngredientService {
	return &ingredientService{repo}
}

func (s *ingredientService) CreateIngredient(ingredient *models.Ingredient) error {
	return s.repo.CreateIngredient(ingredient)
}

func (s *ingredientService) GetIngredients(offset int, limit int) ([]models.Ingredient, error) {
	return s.repo.GetIngredients(offset, limit)
}

func (s *ingredientService) GetIngredientCount() (int64, error) {
	return s.repo.GetIngredientCount()
}

func (s *ingredientService) GetIngredientByID(uuid string) (*models.Ingredient, error) {
	return s.repo.GetIngredientByID(uuid)
}

func (s *ingredientService) UpdateIngredient(ingredient *models.Ingredient) error {
	return s.repo.UpdateIngredient(ingredient)
}

func (s *ingredientService) DeleteIngredient(id uuid.UUID) error {
	return s.repo.DeleteIngredient(id)
}
