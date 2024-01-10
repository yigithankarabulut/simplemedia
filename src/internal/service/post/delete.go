package postservice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
)

func (s *postService) Delete(ctx context.Context, req dto.DeletePostRequest) error {
	var (
		post *model.Post
		err  error
	)
	post, err = s.postStorage.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if post.UserID != req.UserID {
		return errors.New("You are not authorized to delete this post")
	}
	if err := s.postStorage.Delete(ctx, req.ID); err != nil {
		return err
	}
	return nil
}
