package models

import "github.com/google/uuid"

type NotificationSubscription struct {
	BaseModel
	UserID      uuid.UUID         `gorm:"type:uuid;not null" json:"userID"`
	Method      NotifcationMethod `gorm:"type:varchar(20);not null" json:"method"`
	SubjectType SubjectType       `gorm:"type:varchar(20);not null" json:"subjectType"`
}
