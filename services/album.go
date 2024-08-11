package services

import (
	"apcore/models"
	"apcore/repositories"

	"github.com/google/uuid"
)

type AlbumService interface {
	GetAlbum(offset int, limit int, ownerID uuid.UUID) ([]models.CustomerAlbum, error)
	GetAlbumCount(ownerID uuid.UUID) (int64, error)
	AddToAlbum(album *models.CustomerAlbum) error
	DeleteFromAlbum(imageName string, ownerID uuid.UUID) error
}

type albumService struct {
	repo repositories.AlbumRepository
}

func NewAlbumService(repo repositories.AlbumRepository) AlbumService {
	return &albumService{repo}
}

func (s *albumService) GetAlbum(offset int, limit int, ownerID uuid.UUID) ([]models.CustomerAlbum, error) {
	return s.repo.GetAlbum(offset, limit, ownerID)
}

func (s *albumService) GetAlbumCount(ownerID uuid.UUID) (int64, error) {
	return s.repo.GetAlbumCount(ownerID)
}

func (s *albumService) AddToAlbum(album *models.CustomerAlbum) error {
	return s.repo.AddToAlbum(album)
}

func (s *albumService) DeleteFromAlbum(imageName string, ownerID uuid.UUID) error {
	return s.repo.DeleteFromAlbum(imageName, ownerID)
}
