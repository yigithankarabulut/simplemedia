package friendsstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *friendStorage) Insert(ctx context.Context, friend *model.Friends) error {
	if err := r.db.WithContext(ctx).Create(friend).Error; err != nil {
		return err
	}
	return nil
}
