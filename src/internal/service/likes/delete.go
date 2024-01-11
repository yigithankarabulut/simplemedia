package likesservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/likes/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *likesService) Delete(ctx context.Context, req dto.DeleteLikesRequest) error {
	var (
		like model.Likes
	)
	like, err := s.likesStorage.Get(ctx, req.LikeID)
	if err != nil {
		return errors.New(constant.FailedLikeNotFound)
	}
	if like.UserID != req.UserID {
		return errors.New(constant.FailedNotBelong)
	}
	if err := s.likesStorage.Delete(ctx, req.LikeID); err != nil {
		return errors.New(fmt.Sprintf(constant.FailedLikeDelete, err.Error()))
	}
	return nil
}
