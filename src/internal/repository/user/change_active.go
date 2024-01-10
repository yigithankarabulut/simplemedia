package userstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (u *userStorage) ChangeActive(ctx context.Context, id uint, is_active bool) error {
	var user model.User
	if err := u.db.Model(&user).Where("id = ?", id).Update("is_active", is_active).Error; err != nil {
		return err
	}
	return nil
}
