package postservice

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
)

func (s *postService) Create(ctx context.Context, req dto.CreatePostRequest) error {
	var (
		post model.Post
		err  error
		tx   = s.postStorage.CreateTx()
	)
	defer tx.Rollback()
	post.Mapper(req)

	if err = s.postStorage.Insert(ctx, &post, tx); err != nil {
		return err
	}
	tx.Commit()
	return nil
}
