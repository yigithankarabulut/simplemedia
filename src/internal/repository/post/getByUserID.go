package poststorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

func (r *postStorage) GetByUserID(ctx context.Context, userID uint, tx ...*gorm.DB) ([]model.Post, error) {
	db := r.SetTx(tx...)
	var posts []model.Post
	if err := db.WithContext(ctx).Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
