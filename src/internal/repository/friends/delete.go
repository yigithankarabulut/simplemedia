package friendsstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *friendStorage) Delete(ctx context.Context, sender_id, receiver_id uint) error {
	var fr model.Friends
	if err := r.db.Where("sender_id = ? AND receiver_id = ?", sender_id, receiver_id).First(&fr).Error; err != nil {
		return err
	}
	if err := r.db.WithContext(ctx).Delete(&fr, "sender_id = ? AND receiver_id = ?", sender_id, receiver_id).Error; err != nil {
		return err
	}
	return nil
}
