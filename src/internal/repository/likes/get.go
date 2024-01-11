package likesstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *likeStorage) Get(ctx context.Context, id uint) (model.Likes, error) {
	var like model.Likes
	if err := r.db.First(&like, id).Error; err != nil {
		return like, err
	}
	return like, nil
}

func (r *likeStorage) GetComment(ctx context.Context, user_id, id uint) (model.Likes, error) {
	var like model.Likes
	if err := r.db.Where("user_id = ? AND comment_id = ?", user_id, id).First(&like).Error; err != nil {
		return like, err
	}
	return like, nil
}

func (r *likeStorage) GetPost(ctx context.Context, user_id, id uint) (model.Likes, error) {
	var like model.Likes
	if err := r.db.Where("user_id = ? AND post_id = ?", user_id, id).First(&like).Error; err != nil {
		return like, err
	}
	return like, nil
}
