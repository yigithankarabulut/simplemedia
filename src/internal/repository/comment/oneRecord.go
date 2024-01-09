package commentstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *commentStorage) OneRecord(ctx context.Context, condition ...interface{}) (model.Comment, error) {
	var comment model.Comment

	db := r.db.WithContext(ctx)
	if len(condition) >= 2 {
		db = db.Where(condition[0], condition[1:])
	}
	err := db.Debug().First(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}
