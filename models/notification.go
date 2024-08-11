package models

import "github.com/google/uuid"

type NotifcationMethod string

const (
	NotifcationMethodAll   NotifcationMethod = "all"
	NotifcationMethodSMS   NotifcationMethod = "sms"
	NotifcationMethodPush  NotifcationMethod = "push"
	NotifcationMethodInApp NotifcationMethod = "in_app"
)

type Notification struct {
	BaseModel
	Recipient     uuid.UUID         `gorm:"type:uuid;not null" json:"recipient"`
	RecipientUser User              `gorm:"foreignKey:Recipient" json:"recipient_user"`
	Author        string            `gorm:"type:varchar(20);not null" json:"author"`
	Message       string            `gorm:"not null" json:"message"`
	Method        NotifcationMethod `gorm:"type:varchar(20);not null" json:"method"`
	IsRead        bool              `gorm:"not null;default:false" json:"isRead"`
}
