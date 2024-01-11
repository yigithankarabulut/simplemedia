package postservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
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
		return errors.New(fmt.Sprintf(constant.FailedPostCreate, err.Error()))
	}
	tx.Commit()
	return nil
}
