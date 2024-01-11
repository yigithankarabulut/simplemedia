package commentstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

func (r *commentStorage) Update(ctx context.Context, comment *model.Comment, tx ...*gorm.DB) error {
	db := r.SetTx(tx...)
	if err := db.Model(&model.Comment{}).Where("id = ?", comment.ID).Updates(comment).Error; err != nil {
		return err
	}
	return nil
}
