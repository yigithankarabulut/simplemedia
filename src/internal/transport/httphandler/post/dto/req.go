package dto

import "mime/multipart"

type CreatePostRequest struct {
	UserID   uint                  `json:"-"`
	Title    string                `json:"title" form:"title" validate:"required"`
	Content  string                `json:"content" form:"content" validate:"required"`
	Image    *multipart.FileHeader `form:"image" validate:"omitempty,image"`
	ImageUrl string
}

type DeletePostRequest struct {
	UserID uint `json:"-" validate:"-"`
	ID     uint `json:"-" query:"id" param:"id" validate:"required"`
}

type UpdatePostRequest struct {
	UserID  uint   `json:"-" validate:"-"`
	ID      uint   `json:"-" query:"id" param:"id" validate:"required"`
	Title   string `json:"title" form:"title" validate:"omitempty"`
	Content string `json:"content" form:"content" validate:"omitempty"`
	//Image    *multipart.FileHeader `form:"image" validate:"omitempty,image"`
	//ImageUrl string
}

type GetPostRequest struct {
	UserID uint `json:"-" validate:"-"`
	ID     uint `json:"-" query:"id" param:"id" validate:"required"`
}
