package userstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *userStorage) OneRecord(ctx context.Context, condition ...interface{}) (model.User, error) {
	var user model.User

	db := r.db.WithContext(ctx)
	if len(condition) >= 2 {
		db = db.Where(condition[0], condition[1:])
	}
	err := db.Debug().First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
