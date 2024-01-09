package likesstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *likeStorage) GetAllPost(ctx context.Context, id uint) ([]model.Likes, error) {
	var likes []model.Likes
	if err := r.db.WithContext(ctx).Where("post_id = ?", id).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}

func (r *likeStorage) GetAllComment(ctx context.Context, id uint) ([]model.Likes, error) {
	var likes []model.Likes
	if err := r.db.WithContext(ctx).Where("comment_id = ?", id).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}
