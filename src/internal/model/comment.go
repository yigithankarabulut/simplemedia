package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	User     User   `gorm:"foreignKey:UserID"`
	PostID   uint   `json:"post_id"`
	Post     Post   `gorm:"foreignKey:PostID"`
	ParentID uint   `json:"parent_id"`
	Content  string `json:"content"`
	Image    string `json:"image"`
}
