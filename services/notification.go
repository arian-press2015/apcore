package services

import (
	"apcore/models"
	"apcore/repositories"

	"github.com/google/uuid"
)

type NotificationService interface {
	CreateNotification(notification *models.Notification) error
	GetNotifications(offset int, limit int) ([]models.Notification, error)
	GetNotificationCount() (int64, error)
	DeleteNotification(id uuid.UUID) error
	MarkAsRead(id uuid.UUID) error
	MarkAllAsRead() error
}

type notificationService struct {
	repo repositories.NotificationRepository
}

func NewNotificationService(repo repositories.NotificationRepository) NotificationService {
	return &notificationService{repo}
}

func (s *notificationService) CreateNotification(notification *models.Notification) error {
	return s.repo.CreateNotification(notification)
}

func (s *notificationService) GetNotifications(offset int, limit int) ([]models.Notification, error) {
	return s.repo.GetNotifications(offset, limit)
}

func (s *notificationService) GetNotificationCount() (int64, error) {
	return s.repo.GetNotificationCount()
}

func (s *notificationService) DeleteNotification(id uuid.UUID) error {
	return s.repo.DeleteNotification(id)
}

func (s *notificationService) MarkAsRead(id uuid.UUID) error {
	return s.repo.MarkAsRead(id)
}

func (s *notificationService) MarkAllAsRead() error {
	return s.repo.MarkAllAsRead()
}
