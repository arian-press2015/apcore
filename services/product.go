package services

import (
	"apcore/models"
	"apcore/repositories"

	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProducts(offset int, limit int) ([]models.Product, error)
	GetProductCount() (int64, error)
	GetProductByID(uuid string) (*models.Product, error)
	GetProductBySlug(slug string, customerID uuid.UUID) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uuid.UUID) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *productService) GetProducts(offset int, limit int) ([]models.Product, error) {
	return s.repo.GetProducts(offset, limit)
}

func (s *productService) GetProductCount() (int64, error) {
	return s.repo.GetProductCount()
}

func (s *productService) GetProductByID(uuid string) (*models.Product, error) {
	return s.repo.GetProductByID(uuid)
}

func (s *productService) GetProductBySlug(slug string, customerID uuid.UUID) (*models.Product, error) {
	return s.repo.GetProductBySlug(slug, customerID)
}

func (s *productService) UpdateProduct(product *models.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(id uuid.UUID) error {
	return s.repo.DeleteProduct(id)
}
