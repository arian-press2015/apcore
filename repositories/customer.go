package repositories

import (
	"apcore/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomers(offset int, limit int) ([]models.Customer, error)
	GetCustomerByID(id uint) (*models.Customer, error)
	GetCustomerByName(name string) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer) error
	DeleteCustomer(id uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository (db *gorm.DB) CustomerRepository {
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

func (r *customerRepository) GetCustomerByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) GetCustomerByName(name string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("name = ?", name).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) UpdateCustomer(customer *models.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) DeleteCustomer(id uint) error {
	return r.db.Delete(&models.Customer{}, id).Error
}