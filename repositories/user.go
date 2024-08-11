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
