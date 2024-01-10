package dto

type AddFriendRequest struct {
	SenderID   uint `json:"sender_id"`
	ReceiverID uint `json:"receiver_id" query:"receiver_id" validate:"required,min=1,numeric"`
}
