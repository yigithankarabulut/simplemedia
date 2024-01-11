package commentstorage

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
)

func (r *commentStorage) GetAllByCommentID(ctx context.Context, parentID uint) ([]model.Comment, error) {
	var (
		comments []model.Comment
		err      error
	)
	if err = r.db.Where("parent_id = ?", parentID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentStorage) GetAllByPostID(ctx context.Context, postID uint) ([]model.Comment, error) {
	var (
		comments []model.Comment
		err      error
	)
	if err = r.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
