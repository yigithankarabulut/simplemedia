package likesstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

type LikeStorer interface {
	Insert(ctx context.Context, like *model.Likes) error
	Delete(ctx context.Context, id uint) error
	PostDelete(ctx context.Context, id uint) error
	CommentDelete(ctx context.Context, id uint) error
	Get(ctx context.Context, id uint) (model.Likes, error)
	GetComment(ctx context.Context, user_id, id uint) (model.Likes, error)
	GetPost(ctx context.Context, user_id, id uint) (model.Likes, error)
	GetAllPost(ctx context.Context, id uint) ([]model.Likes, error)
	GetAllComment(ctx context.Context, id uint) ([]model.Likes, error)
}

type likeStorage struct {
	db *gorm.DB
}

type LikeRepositoryOption func(*likeStorage)

func WithLikeRepositoryDB(db *gorm.DB) LikeRepositoryOption {
	return func(u *likeStorage) {
		u.db = db
	}
}

func NewLikeRepository(opts ...LikeRepositoryOption) LikeStorer {
	u := &likeStorage{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
