package services

import (
	"apcore/models"
	"apcore/repositories"

	"github.com/google/uuid"
)

type CustomerService interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomers(offset int, limit int) ([]models.Customer, error)
	GetCustomerCount() (int64, error)
	GetCustomerByID(uuid string) (*models.Customer, error)
	GetCustomerBySlug(slug string) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer) error
	DeleteCustomer(id uuid.UUID) error
	GetAlbum(offset int, limit int, ownerID uuid.UUID) ([]models.CustomerAlbum, error)
	GetAlbumCount() (int64, error)
	AddToAlbum(album *models.CustomerAlbum) error
	DeleteFromAlbum(imageName string, ownerID uuid.UUID) error
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

func (s *customerService) GetCustomerCount() (int64, error) {
	return s.repo.GetCustomerCount()
}

func (s *customerService) GetCustomerByID(uuid string) (*models.Customer, error) {
	return s.repo.GetCustomerByID(uuid)
}

func (s *customerService) GetCustomerBySlug(slug string) (*models.Customer, error) {
	return s.repo.GetCustomerBySlug(slug)
}

func (s *customerService) UpdateCustomer(customer *models.Customer) error {
	return s.repo.UpdateCustomer(customer)
}

func (s *customerService) DeleteCustomer(id uuid.UUID) error {
	return s.repo.DeleteCustomer(id)
}

func (s *customerService) GetAlbum(offset int, limit int, ownerID uuid.UUID) ([]models.CustomerAlbum, error) {
	return s.repo.GetAlbum(offset, limit, ownerID)
}

func (s *customerService) GetAlbumCount() (int64, error) {
	return s.repo.GetAlbumCount()
}

func (s *customerService) AddToAlbum(album *models.CustomerAlbum) error {
	return s.repo.AddToAlbum(album)
}

func (s *customerService) DeleteFromAlbum(imageName string, ownerID uuid.UUID) error {
	return s.repo.DeleteFromAlbum(imageName, ownerID)
}
