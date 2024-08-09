package repositories

import (
	"apcore/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomers(offset int, limit int) ([]models.Customer, error)
	GetCustomerCount() (int64, error)
	GetCustomerByID(uuid string) (*models.Customer, error)
	GetCustomerBySlug(slug string) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer) error
	DeleteCustomer(id uuid.UUID) error
	CheckUserHasAccessToCustomer(userID uuid.UUID, customerID uuid.UUID) (bool, error)
	GetAlbum(offset int, limit int, ownerID uuid.UUID) ([]models.CustomerAlbum, error)
	GetAlbumCount() (int64, error)
	AddToAlbum(album *models.CustomerAlbum) error
	DeleteFromAlbum(imageName string, ownerID uuid.UUID) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) CreateCustomer(customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) GetCustomers(offset int, limit int) ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.Offset(offset).Limit(limit).Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *customerRepository) GetCustomerCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.Customer{}).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *customerRepository) GetCustomerByID(uuid string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("id = ?", uuid).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) GetCustomerBySlug(slug string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("slug = ?", slug).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) UpdateCustomer(customer *models.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) DeleteCustomer(id uuid.UUID) error {
	return r.db.Delete(&models.Customer{}, id).Error
}

func (r *customerRepository) CheckUserHasAccessToCustomer(userID uuid.UUID, customerID uuid.UUID) (bool, error) {
	var userCustomer models.UserCustomer
	err := r.db.Where("user_id = ? AND customer_id = ?", userID, customerID).First(&userCustomer).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *customerRepository) GetAlbum(offset int, limit int, ownerID uuid.UUID) ([]models.CustomerAlbum, error) {
	var album []models.CustomerAlbum
	err := r.db.Where("owner_id = ?", ownerID).Offset(offset).Limit(limit).Find(&album).Error
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (r *customerRepository) GetAlbumCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.CustomerAlbum{}).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *customerRepository) AddToAlbum(album *models.CustomerAlbum) error {
	return r.db.Create(album).Error
}

func (r *customerRepository) DeleteFromAlbum(name string, ownerID uuid.UUID) error {
	return r.db.Where("name = ? AND owner_id = ?", name, ownerID).Delete(&models.CustomerAlbum{}).Error
}
