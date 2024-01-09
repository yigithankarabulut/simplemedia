package userstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

func (r *userStorage) Insert(ctx context.Context, user *model.User, tx ...*gorm.DB) error {
	db := r.SetTx(tx...)
	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}
