package dto

type ResponseForGetPost struct {
	UserID   uint   `json:"user_id"`
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageUrl string `json:"image"`
	ShortUrl string `json:"short_url"`
}

type ResponseForGetAllPost struct {
	UserID   uint   `json:"user_id"`
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageUrl string `json:"image"`
}
