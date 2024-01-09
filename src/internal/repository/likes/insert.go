package likesstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *likeStorage) Insert(ctx context.Context, like *model.Likes) error {
	if err := r.db.WithContext(ctx).Create(like).Error; err != nil {
		return err
	}
	return nil
}
