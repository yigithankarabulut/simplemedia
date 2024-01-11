package model

import "gorm.io/gorm"

type Friends struct {
	gorm.Model
	SenderID   uint `json:"sender_id"`
	ReceiverID uint `json:"receiver_id"`
	Accepted   bool `json:"accepted"`
}
