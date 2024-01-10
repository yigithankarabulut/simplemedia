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

func (r *userStorage) OneRecordWithPhone(ctx context.Context, phone string) (model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("phone=?", phone).Debug().First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userStorage) OneRecordWithEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("email=?", email).Debug().First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userStorage) OneRecordWithUserName(ctx context.Context, userName string) (model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("username=?", userName).Debug().First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
