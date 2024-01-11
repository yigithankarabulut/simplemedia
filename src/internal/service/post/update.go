package postservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
	"time"
)

func (s *postService) Update(ctx context.Context, req dto.UpdatePostRequest) error {
	var (
		err  error
		post *model.Post
	)
	post, err = s.postStorage.GetByID(ctx, req.ID)
	if err != nil {
		return errors.New(fmt.Sprintf(constant.IDNotFound, err.Error()))
	}
	if time.Now().Sub(post.CreatedAt).Minutes() > 5 {
		return errors.New(fmt.Sprintf(constant.FailedPostUpdate, "You can update the post within 5 minutes after creating it."))
	}
	if err = s.postStorage.Update(ctx, post); err != nil {
		return errors.New(fmt.Sprintf(constant.FailedPostUpdate, err.Error()))
	}
	return nil
}
