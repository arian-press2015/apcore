package repositories

import (
	"apcore/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUsers(offset int, limit int) ([]models.User, error)
	GetUserByID(uuid uuid.UUID) (*models.User, error)
	GetUserByPhone(username string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uuid.UUID) error
	GetFavorites(offset int, limit int, userID uuid.UUID) ([]models.Favorites, error)
	GetFavoritesCount(userID uuid.UUID) (int64, error)
	AddToFavorites(favorite *models.Favorites) error
	DeleteFromFavorites(customerID uuid.UUID, userID uuid.UUID) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUsers(offset int, limit int) ([]models.User, error) {
	var users []models.User
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(uuid uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", uuid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id uuid.UUID) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) GetFavorites(offset int, limit int, userID uuid.UUID) ([]models.Favorites, error) {
	var favorites []models.Favorites
	err := r.db.Where("user_id = ?", userID).Offset(offset).Limit(limit).Preload("Customer").Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (r *userRepository) GetFavoritesCount(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&models.Favorites{}).Where("user_id = ?", userID).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *userRepository) AddToFavorites(favorite *models.Favorites) error {
	return r.db.Create(favorite).Error
}

func (r *userRepository) DeleteFromFavorites(customerID uuid.UUID, userID uuid.UUID) error {
	return r.db.Unscoped().Where("customer_id = ? AND user_id = ?", customerID, userID).Delete(&models.Favorites{}).Error
}
