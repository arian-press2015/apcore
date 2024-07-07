package models

type User struct {
	BaseModel
	FullName      string                   `gorm:"not null" json:"fullName" binding:"required"`
	Phone         string                   `gorm:"unique;not null;size:11" json:"phone" binding:"required"`
	ProfileImage  string                   `json:"profile_image"`
	Nid           string                   `gorm:"size:10;unique" json:"nid"`
	Verified      bool                     `gorm:"default:false" json:"verified"`
	Roles         []Role                   `gorm:"many2many:user_roles;" json:"roles"`
	Customers     []Customer               `gorm:"many2many:user_customers;" json:"customers"`
	Notifications []Notification           `gorm:"foreignKey:Recipient" json:"notifications"`
	Subscription  NotificationSubscription `json:"subscription"`
}
