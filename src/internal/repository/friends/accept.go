package friendsstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *friendStorage) Accept(ctx context.Context, sender_id, receiver_id uint) error {
	if err := r.db.WithContext(ctx).Model(&model.Friends{}).Where("sender_id = ? AND receiver_id = ?", sender_id, receiver_id).Update("accepted", true).Error; err != nil {
		return err
	}
	return nil
}
