package friendsstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

type FriendStorer interface {
	Insert(ctx context.Context, friend *model.Friends) error
	Delete(ctx context.Context, id uint) error
	Accept(ctx context.Context, sender_id, receiver_id uint) error
	Reject(ctx context.Context, sender_id, receiver_id uint) error
}

type friendStorage struct {
	db *gorm.DB
}

type FriendRepositoryOption func(*friendStorage)

func WithFriendRepositoryDB(db *gorm.DB) FriendRepositoryOption {
	return func(u *friendStorage) {
		u.db = db
	}
}

func NewFriendRepository(opts ...FriendRepositoryOption) FriendStorer {
	u := &friendStorage{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
