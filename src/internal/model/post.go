package model

import (
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func (p *Post) Mapper(req dto.CreatePostRequest) {
	p.UserID = req.UserID
	p.Title = req.Title
	p.Content = req.Content
	p.Image = req.ImageUrl
}
