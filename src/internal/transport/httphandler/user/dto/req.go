package dto

type RegisterRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
	Phone    string `json:"phone" validate:"required,e164"`
	Picture  string `json:"picture" validate:"omitempty,image"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type ChangePwdRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=8,max=50"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=50"`
}
