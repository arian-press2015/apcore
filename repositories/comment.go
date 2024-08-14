package repositories

import (
	"apcore/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) error
	GetComments(offset int, limit int) ([]models.Comment, error)
	GetCommentCount() (int64, error)
	GetCommentByID(uuid string) (*models.Comment, error)
	UpdateComment(comment *models.Comment) error
	DeleteComment(id uuid.UUID) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) CreateComment(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) GetComments(offset int, limit int) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Offset(offset).Limit(limit).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepository) GetCommentCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.Comment{}).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *commentRepository) GetCommentByID(uuid string) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.Where("id = ?", uuid).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) UpdateComment(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) DeleteComment(id uuid.UUID) error {
	return r.db.Delete(&models.Comment{}, id).Error
}
