package poststorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

type PostStorer interface {
	Insert(ctx context.Context, post *model.Post, tx ...*gorm.DB) error
	OneRecord(ctx context.Context, condition ...interface{}) (model.Post, error)
	GetByID(ctx context.Context, id uint, tx ...*gorm.DB) (*model.Post, error)
	GetByUserID(ctx context.Context, userID uint, tx ...*gorm.DB) ([]model.Post, error)
	Update(ctx context.Context, post *model.Post, tx ...*gorm.DB) error
	Delete(ctx context.Context, id uint, tx ...*gorm.DB) error
	CreateTx() *gorm.DB
	CommitTx(tx *gorm.DB)
	RollBackTx(tx *gorm.DB)
}

type postStorage struct {
	db *gorm.DB
}

type PostRepositoryOption func(*postStorage)

func WithPostRepositoryDB(db *gorm.DB) PostRepositoryOption {
	return func(u *postStorage) {
		u.db = db
	}
}

func NewPostRepository(opts ...PostRepositoryOption) PostStorer {
	u := &postStorage{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
