package dto

type BaseLikesResponse struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
}

type GetAllLikesResponse struct {
	Likes []BaseLikesResponse `json:"likes"`
	Count uint                `json:"count"`
}
