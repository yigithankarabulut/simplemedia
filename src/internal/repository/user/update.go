package userstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

func (r *userStorage) Update(ctx context.Context, user *model.User, tx ...*gorm.DB) error {
	db := r.SetTx(tx...)
	if err := db.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}
