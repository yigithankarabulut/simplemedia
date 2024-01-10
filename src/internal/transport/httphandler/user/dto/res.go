package dto

type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Token    string `json:"token"`
}
