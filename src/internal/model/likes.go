package model

import "gorm.io/gorm"

type Likes struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	PostID    uint `json:"post_id"`
	CommentID uint `json:"comment_id"`
}
