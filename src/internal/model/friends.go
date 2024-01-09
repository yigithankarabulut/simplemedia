package model

type Friends struct {
	SenderID   uint `json:"sender_id"`
	ReceiverID uint `json:"receiver_id"`
	Accepted   bool `json:"accepted"`
}
