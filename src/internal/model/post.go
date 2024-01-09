package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}
