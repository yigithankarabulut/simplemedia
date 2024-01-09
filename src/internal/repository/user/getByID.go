package userstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

func (r *userStorage) GetByID(ctx context.Context, id uint, tx ...*gorm.DB) (*model.User, error) {
	db := r.SetTx(tx...)
	var user model.User
	if err := db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
