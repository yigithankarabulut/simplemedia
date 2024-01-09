package friendsstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *friendStorage) Reject(ctx context.Context, sender_id, receiver_id uint) error {
	if err := r.db.WithContext(ctx).Delete(&model.Friends{}, "sender_id = ? AND receiver_id = ?", sender_id, receiver_id).Error; err != nil {
		return err
	}
	return nil
}
