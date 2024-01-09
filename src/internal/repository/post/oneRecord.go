package poststorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *postStorage) OneRecord(ctx context.Context, conditions ...interface{}) (model.Post, error) {
	var post model.Post

	if len(conditions) >= 2 {
		r.db = r.db.Where(conditions[0], conditions[1:])
	}
	err := r.db.Debug().First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}
