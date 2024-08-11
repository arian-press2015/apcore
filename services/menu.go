package services

import (
	"apcore/dto"
	"apcore/models"
	"apcore/repositories"
	"errors"

	"github.com/google/uuid"
)

type MenuService interface {
	GetMenu(customerID uuid.UUID) (*models.Menu, error)
	CreateMenu(customerID uuid.UUID, dto dto.CreateMenuBody) error
	UpdateMenu(customerID uuid.UUID, dto dto.UpdateMenuBody) error
}

type menuService struct {
	repo repositories.MenuRepository
}

func NewMenuService(repo repositories.MenuRepository) MenuService {
	return &menuService{repo}
}

func (s *menuService) GetMenu(customerID uuid.UUID) (*models.Menu, error) {
	return s.repo.GetMenu(customerID)
}

func (s *menuService) CreateMenu(customerID uuid.UUID, dto dto.CreateMenuBody) error {
	categories, err := s.constrcutCategories(dto.Categories)
	if err != nil {
		return errors.New("invalid source categories")
	}

	menu := &models.Menu{
		CustomerID: customerID,
		Categories: categories,
	}

	return s.repo.CreateMenu(menu)
}

func (s *menuService) UpdateMenu(customerID uuid.UUID, dto dto.UpdateMenuBody) error {
	menu, err := s.repo.GetMenu(customerID)
	if err != nil {
		return errors.New("invalid customerID")
	}

	err = s.desctructCategories(menu.Categories)
	if err != nil {
		return errors.New("error while removing orphan categories")
	}

	categories, err := s.constrcutCategories(dto.Categories)
	if err != nil {
		return errors.New("invalid source categories")
	}

	menu.Categories = categories

	return s.repo.UpdateMenu(menu)
}

func (s *menuService) constrcutCategories(categoriesString []string) ([]models.Category, error) {
	categories := []models.Category{}

	// construct categories slice e.g.
	// for i := 1; i <= 5; i++ {
	//     category := models.Category{
	//         Name:        fmt.Sprintf("Category %d", i),
	//         Description: fmt.Sprintf("Description for category %d", i),
	//         ID:          uuid.New(), // Example of adding a unique ID
	//     }

	//     // Append the calculated category to the slice
	//     categories = append(categories, category)
	// }

	return categories, nil
}

func (s *menuService) desctructCategories(categories []models.Category) error {
	// do sth like removing current categories so we don't have orphan entries
	return nil
}
