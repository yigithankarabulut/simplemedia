package userstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

type UserStorer interface {
	Insert(ctx context.Context, user *model.User, tx ...*gorm.DB) error
	GetByID(ctx context.Context, id uint, tx ...*gorm.DB) (*model.User, error)
	GetByUsername(ctx context.Context, username string, tx ...*gorm.DB) (*model.User, error)
	Update(ctx context.Context, user *model.User, tx ...*gorm.DB) error
	Delete(ctx context.Context, id uint, tx ...*gorm.DB) error
	OneRecord(ctx context.Context, condition ...interface{}) (model.User, error)
	OneRecordWithPhone(ctx context.Context, phone string) (model.User, error)
	OneRecordWithEmail(ctx context.Context, email string) (model.User, error)
	OneRecordWithUserName(ctx context.Context, userName string) (model.User, error)
	ChangeActive(ctx context.Context, id uint, is_active bool) error
	CreateTx() *gorm.DB
	CommitTx(tx *gorm.DB)
	RollBackTx(tx *gorm.DB)
}

type userStorage struct {
	db *gorm.DB
}

type UserRepositoryOption func(*userStorage)

func WithUserRepositoryDB(db *gorm.DB) UserRepositoryOption {
	return func(u *userStorage) {
		u.db = db
	}
}

func NewUserRepository(opts ...UserRepositoryOption) UserStorer {
	u := &userStorage{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
