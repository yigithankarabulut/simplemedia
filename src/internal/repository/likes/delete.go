package likesstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *likeStorage) PostDelete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&model.Likes{}, "post_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *likeStorage) CommentDelete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&model.Likes{}, "comment_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
