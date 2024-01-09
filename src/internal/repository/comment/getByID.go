package commentstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

func (r *commentStorage) GetByID(ctx context.Context, id uint, tx ...*gorm.DB) (*model.Comment, error) {
	var comment model.Comment

	db := r.SetTx(tx...)
	if err := db.WithContext(ctx).Where("id = ?", id).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}
