package model

import "gorm.io/gorm"

type Likes struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	User      User `gorm:"foreignKey:UserID"`
	PostID    uint `json:"post_id"`
	CommentID uint `json:"comment_id"`
}
