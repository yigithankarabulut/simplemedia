package poststorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

func (r *postStorage) GetByID(ctx context.Context, id uint, tx ...*gorm.DB) (*model.Post, error) {
	db := r.SetTx(tx...)
	var post model.Post
	if err := db.WithContext(ctx).Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
