package poststorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/gorm"
)

func (r *postStorage) Delete(ctx context.Context, id uint, tx ...*gorm.DB) error {
	db := r.SetTx(tx...)
	if err := db.WithContext(ctx).Delete(&model.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}
