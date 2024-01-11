package dto

import (
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"mime/multipart"
)

type CreateCommentRequest struct {
	UserID   uint                  `json:"-"`
	PostID   uint                  `json:"post_id" form:"post_id,type=text" validate:"required"`
	ParentID uint                  `json:"parent_id" form:"parent_id,type=text" validate:"omitempty"`
	Content  string                `json:"content" form:"content" validate:"required"`
	Image    *multipart.FileHeader `form:"image" validate:"omitempty,image"`
	ImageUrl string
}

func (r CreateCommentRequest) Convert() *model.Comment {
	return &model.Comment{
		UserID:   r.UserID,
		PostID:   r.PostID,
		ParentID: r.ParentID,
		Content:  r.Content,
		Image:    r.ImageUrl,
	}
}

type UpdateCommentRequest struct {
	UserID  uint   `json:"-"`
	ID      uint   `json:"post_id" form:"post_id,type=text" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
}

type DeleteCommentRequest struct {
	UserID uint `json:"-"`
	ID     uint `json:"-" query:"post_id" validate:"required"`
}

type GetCommentRequest struct {
	UserID uint `json:"-"`
	ID     uint `json:"-" query:"post_id" validate:"required"`
}

type GetAllCommentRequest struct {
	UserID    uint `json:"-"`
	PostID    uint `json:"-" query:"post_id" validate:"omitempty,numeric"`
	CommentID uint `json:"-" query:"comment_id" validate:"omitempty,numeric"`
}
