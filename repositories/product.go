package repositories

import (
	"apcore/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProducts(offset int, limit int) ([]models.Product, error)
	GetProductCount() (int64, error)
	GetProductByID(uuid string) (*models.Product, error)
	GetProductBySlug(slug string, customerID uuid.UUID) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uuid.UUID) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetProducts(offset int, limit int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetProductCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.Product{}).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *productRepository) GetProductByID(uuid string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ?", uuid).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetProductBySlug(slug string, customerID uuid.UUID) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("slug = ? AND customer_id = ?", slug, customerID).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) DeleteProduct(id uuid.UUID) error {
	return r.db.Delete(&models.Product{}, id).Error
}
