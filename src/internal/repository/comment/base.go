package commentstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

type CommentStorer interface {
	Insert(ctx context.Context, comment *model.Comment, tx ...*gorm.DB) error
	OneRecord(ctx context.Context, condition ...interface{}) (model.Comment, error)
	GetByID(ctx context.Context, id uint, tx ...*gorm.DB) (*model.Comment, error)
	GetAllByCommentID(ctx context.Context, commentID uint) ([]model.Comment, error)
	GetAllByPostID(ctx context.Context, postID uint) ([]model.Comment, error)
	Update(ctx context.Context, comment *model.Comment, tx ...*gorm.DB) error
	Delete(ctx context.Context, id uint, tx ...*gorm.DB) error
	CreateTx() *gorm.DB
	CommitTx(tx *gorm.DB)
	RollBackTx(tx *gorm.DB)
}

type commentStorage struct {
	db *gorm.DB
}

type CommentRepositoryOption func(*commentStorage)

func WithCommentRepositoryDB(db *gorm.DB) CommentRepositoryOption {
	return func(u *commentStorage) {
		u.db = db
	}
}

func NewCommentRepository(opts ...CommentRepositoryOption) CommentStorer {
	u := &commentStorage{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
