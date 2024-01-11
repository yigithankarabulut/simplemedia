package dto

import "github.com/yigithankarabulut/simplemedia/src/internal/model"

type BaseCommentResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	ParentID uint   `json:"parent_id"`
	Content  string `json:"content"`
	Image    string `json:"image"`
}

func (r *BaseCommentResponse) Convert(comment *model.Comment) {
	r.ID = comment.ID
	r.Username = comment.User.Username
	r.ParentID = comment.ParentID
	r.Content = comment.Content
	r.Image = comment.Image
}
