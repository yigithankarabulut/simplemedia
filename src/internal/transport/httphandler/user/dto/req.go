package dto

import "mime/multipart"

type RegisterRequest struct {
	Username   string                `json:"username" form:"username" validate:"required,alphanum,min=3,max=50"`
	Email      string                `json:"email" form:"email" validate:"required,email"`
	Password   string                `json:"password" form:"password" validate:"required,min=8,max=50"`
	Phone      string                `json:"phone" form:"phone" validate:"required,e164"`
	Picture    *multipart.FileHeader `form:"picture" validate:"omitempty,image"`
	PictureUrl string
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type ChangePwdRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=8,max=50"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=50"`
}

type UpdatePictureRequest struct {
	UserID     uint                  `json:"-"`
	Username   string                `json:"-" form:"username" validate:"required,alphanum,min=3,max=50"`
	Picture    *multipart.FileHeader `json:"-" form:"picture" validate:"omitempty,image"`
	PictureUrl string
}
