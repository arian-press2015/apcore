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
	GetFavorites(offset int, limit int, userID uuid.UUID) ([]models.Favorites, error)
	GetFavoritesCount(userID uuid.UUID) (int64, error)
	AddToFavorites(favorite *models.Favorites) error
	DeleteFromFavorites(customerID uuid.UUID, userID uuid.UUID) error
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

func (s *userService) GetFavorites(offset int, limit int, userID uuid.UUID) ([]models.Favorites, error) {
	return s.repo.GetFavorites(offset, limit, userID)
}

func (s *userService) GetFavoritesCount(userID uuid.UUID) (int64, error) {
	return s.repo.GetFavoritesCount(userID)
}

func (s *userService) AddToFavorites(favorite *models.Favorites) error {
	return s.repo.AddToFavorites(favorite)
}

func (s *userService) DeleteFromFavorites(customerID uuid.UUID, userID uuid.UUID) error {
	return s.repo.DeleteFromFavorites(customerID, userID)
}
