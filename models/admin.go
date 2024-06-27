package models

type Admin struct {
	BaseModel
	Username     string `gorm:"unique;not null" json:"username" binding:"required"`
	Phone        string `gorm:"unique;not null;size:11" json:"phone" binding:"required"`
	ProfileImage string `json:"profile_image"`
	Password     string `gorm:"not null" json:"password" binding:"required"`
	TOTPSecret   string `json:"totp_secret"`
}
