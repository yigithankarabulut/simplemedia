package dto

import "github.com/yigithankarabulut/simplemedia/src/internal/model"

type BaseLikesRequest struct {
	UserID    uint `json:"-"`
	PostID    uint `json:"post_id" query:"post_id" validate:"omitempty"`
	CommentID uint `json:"comment_id" query:"comment_id" validate:"omitempty"`
}

func (r *BaseLikesRequest) Convert() *model.Likes {
	return &model.Likes{
		UserID:    r.UserID,
		PostID:    r.PostID,
		CommentID: r.CommentID,
	}
}

type DeleteLikesRequest struct {
	UserID uint `json:"-"`
	LikeID uint `json:"-" query:"id" validate:"required"`
}
