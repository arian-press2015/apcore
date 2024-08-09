package repositories

import (
	"apcore/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	CreateNotification(notification *models.Notification) error
	GetNotifications(offset int, limit int) ([]models.Notification, error)
	GetNotificationCount() (int64, error)
	DeleteNotification(id uuid.UUID) error
	MarkAsRead(id uuid.UUID) error
	MarkAllAsRead() error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db}
}

func (r *notificationRepository) CreateNotification(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) GetNotifications(offset int, limit int) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.Offset(offset).Limit(limit).Preload("RecipientUser").Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (s *notificationRepository) GetNotificationCount() (int64, error) {
	var count int64
	err := s.db.Model(&models.Notification{}).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *notificationRepository) DeleteNotification(id uuid.UUID) error {
	return r.db.Delete(&models.Notification{}, id).Error
}

func (r *notificationRepository) MarkAsRead(id uuid.UUID) error {
	var record *models.Notification

	if err := r.db.Select("is_read").Where("id = ?", id).First(&record).Error; err != nil {
		return err
	}

	newIsRead := !record.IsRead

	err := r.db.Model(&models.Notification{}).
		Where("id = ?", id).
		Update("is_read", newIsRead).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *notificationRepository) MarkAllAsRead() error {
	err := r.db.Model(&models.Notification{}).
		Where("is_read = ?", false).
		Update("is_read", true).Error

	if err != nil {
		return err
	}

	return nil
}
