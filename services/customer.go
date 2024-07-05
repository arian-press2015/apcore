package services

import (
	"apcore/models"
	"apcore/repositories"
)

type CustomerService interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomers(offset int, limit int) ([]models.Customer, error)
	GetCustomerByID(uuid string) (*models.Customer, error)
	GetCustomerByName(name string) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer) error
	DeleteCustomer(id uint) error
}

type customerService struct {
	repo repositories.CustomerRepository
}

func NewCustomerService(repo repositories.CustomerRepository) CustomerService {
	return &customerService{repo}
}

func (s *customerService) CreateCustomer(customer *models.Customer) error {
	return s.repo.CreateCustomer(customer)
}

func (s *customerService) GetCustomers(offset int, limit int) ([]models.Customer, error) {
	return s.repo.GetCustomers(offset, limit)
}

func (s *customerService) GetCustomerByID(uuid string) (*models.Customer, error) {
	return s.repo.GetCustomerByID(uuid)
}

func (s *customerService) GetCustomerByName(name string) (*models.Customer, error) {
	return s.repo.GetCustomerByName(name)
}

func (s *customerService) UpdateCustomer(customer *models.Customer) error {
	return s.repo.UpdateCustomer(customer)
}

func (s *customerService) DeleteCustomer(id uint) error {
	return s.repo.DeleteCustomer(id)
}
