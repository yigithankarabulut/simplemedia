package friendsstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *friendStorage) Get(ctx context.Context, sender_id, receiver_id uint) (model.Friends, error) {
	var (
		friend model.Friends
	)
	if err := r.db.Where("sender_id = ? AND receiver_id = ?", sender_id, receiver_id).First(&friend).Error; err != nil {
		return friend, err
	}
	return friend, nil
}
