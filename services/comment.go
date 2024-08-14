package services

import (
	"apcore/models"
	"apcore/repositories"

	"github.com/google/uuid"
)

type CommentService interface {
	CreateComment(comment *models.Comment) error
	GetComments(offset int, limit int) ([]models.Comment, error)
	GetCommentCount() (int64, error)
	GetCommentByID(uuid string) (*models.Comment, error)
	UpdateComment(comment *models.Comment) error
	DeleteComment(id uuid.UUID) error
}

type commentService struct {
	repo repositories.CommentRepository
}

func NewCommentService(repo repositories.CommentRepository) CommentService {
	return &commentService{repo}
}

func (s *commentService) CreateComment(comment *models.Comment) error {
	return s.repo.CreateComment(comment)
}

func (s *commentService) GetComments(offset int, limit int) ([]models.Comment, error) {
	return s.repo.GetComments(offset, limit)
}

func (s *commentService) GetCommentCount() (int64, error) {
	return s.repo.GetCommentCount()
}

func (s *commentService) GetCommentByID(uuid string) (*models.Comment, error) {
	return s.repo.GetCommentByID(uuid)
}

func (s *commentService) UpdateComment(comment *models.Comment) error {
	return s.repo.UpdateComment(comment)
}

func (s *commentService) DeleteComment(id uuid.UUID) error {
	return s.repo.DeleteComment(id)
}
