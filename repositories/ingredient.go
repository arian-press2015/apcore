package repositories

import (
	"apcore/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IngredientRepository interface {
	CreateIngredient(ingredient *models.Ingredient) error
	GetIngredients(offset int, limit int) ([]models.Ingredient, error)
	GetIngredientCount() (int64, error)
	GetIngredientByID(uuid string) (*models.Ingredient, error)
	UpdateIngredient(ingredient *models.Ingredient) error
	DeleteIngredient(id uuid.UUID) error
}

type ingredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) IngredientRepository {
	return &ingredientRepository{db}
}

func (r *ingredientRepository) CreateIngredient(ingredient *models.Ingredient) error {
	return r.db.Create(ingredient).Error
}

func (r *ingredientRepository) GetIngredients(offset int, limit int) ([]models.Ingredient, error) {
	var ingredients []models.Ingredient
	err := r.db.Offset(offset).Limit(limit).Find(&ingredients).Error
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

func (r *ingredientRepository) GetIngredientCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.Ingredient{}).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ingredientRepository) GetIngredientByID(uuid string) (*models.Ingredient, error) {
	var ingredient models.Ingredient
	err := r.db.Where("id = ?", uuid).First(&ingredient).Error
	if err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (r *ingredientRepository) UpdateIngredient(ingredient *models.Ingredient) error {
	return r.db.Save(ingredient).Error
}

func (r *ingredientRepository) DeleteIngredient(id uuid.UUID) error {
	return r.db.Delete(&models.Ingredient{}, id).Error
}
