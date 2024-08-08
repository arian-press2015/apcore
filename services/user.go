package services

import (
	"apcore/models"
	"apcore/repositories"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUsers(offset int, limit int) ([]models.User, error)
	GetUserByID(uuid uuid.UUID) (*models.User, error)
	GetUserByPhone(phone string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUsers(offset int, limit int) ([]models.User, error) {
	return s.repo.GetUsers(offset, limit)
}

func (s *userService) GetUserByID(uuid uuid.UUID) (*models.User, error) {
	return s.repo.GetUserByID(uuid)
}

func (s *userService) GetUserByPhone(phone string) (*models.User, error) {
	return s.repo.GetUserByPhone(phone)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
