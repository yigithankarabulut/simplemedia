package poststorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

func (r *postStorage) Update(ctx context.Context, post *model.Post, tx ...*gorm.DB) error {
	db := r.SetTx(tx...)
	if err := db.WithContext(ctx).Save(post).Error; err != nil {
		return err
	}
	return nil
}
