package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string `json:"username" gorm:"size:50"`
	Email          string `json:"email" gorm:"size:50"`
	Password       string `json:"password"`
	Phone          string `json:"phone" gorm:"uniqueIndex;type:varchar(20)"`
	ProfilePicture string `json:"profile_picture"`
	IsActive       bool   `json:"is_active"`
}
