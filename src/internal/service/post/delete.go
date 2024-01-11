package postservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *postService) Delete(ctx context.Context, req dto.DeletePostRequest) error {
	var (
		post *model.Post
		err  error
	)
	post, err = s.postStorage.GetByID(ctx, req.ID)
	if err != nil {
		return errors.New(fmt.Sprintf(constant.IDNotFound, err.Error()))
	}
	if post.UserID != req.UserID {
		return errors.New(constant.FailedDeleteUnauthorized)
	}
	if err := s.postStorage.Delete(ctx, req.ID); err != nil {
		return errors.New(fmt.Sprintf(constant.FailedPostDelete, err.Error()))
	}
	return nil
}
