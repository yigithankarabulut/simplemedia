package commentstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

func (r *commentStorage) Insert(ctx context.Context, comment *model.Comment, tx ...*gorm.DB) error {
	db := r.SetTx(tx...)
	if err := db.WithContext(ctx).Create(comment).Error; err != nil {
		return err
	}
	return nil
}
