package postservice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"time"
)

func (s *postService) Update(ctx context.Context, req dto.UpdatePostRequest) error {
	var (
		err  error
		post *model.Post
	)
	post, err = s.postStorage.GetByID(ctx, req.ID)
	if err != nil {
		return errors.New("post not found")
	}
	if time.Now().Sub(post.CreatedAt).Minutes() > 5 {
		return errors.New("post can not be updated")
	}
	if err = s.postStorage.Update(ctx, post); err != nil {
		return err
	}
	return nil
}
