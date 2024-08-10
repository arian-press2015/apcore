package repositories

import (
	"apcore/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlbumRepository interface {
	GetAlbum(offset int, limit int, ownerID uuid.UUID) ([]models.CustomerAlbum, error)
	GetAlbumCount(ownerID uuid.UUID) (int64, error)
	AddToAlbum(album *models.CustomerAlbum) error
	DeleteFromAlbum(imageName string, ownerID uuid.UUID) error
}

type albumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) AlbumRepository {
	return &albumRepository{db}
}

func (r *albumRepository) GetAlbum(offset int, limit int, ownerID uuid.UUID) ([]models.CustomerAlbum, error) {
	var album []models.CustomerAlbum
	err := r.db.Where("owner_id = ?", ownerID).Offset(offset).Limit(limit).Find(&album).Error
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (r *albumRepository) GetAlbumCount(ownerID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&models.CustomerAlbum{}).Where("owner_id = ?", ownerID).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *albumRepository) AddToAlbum(album *models.CustomerAlbum) error {
	return r.db.Create(album).Error
}

func (r *albumRepository) DeleteFromAlbum(name string, ownerID uuid.UUID) error {
	return r.db.Unscoped().Where("name = ? AND owner_id = ?", name, ownerID).Delete(&models.CustomerAlbum{}).Error
}
