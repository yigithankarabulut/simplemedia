package likesservice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/likes/dto"
)

func (s *likesService) Delete(ctx context.Context, req dto.DeleteLikesRequest) error {
	var (
		like model.Likes
	)
	like, err := s.likesStorage.Get(ctx, req.LikeID)
	if err != nil {
		return errors.New("like not found")
	}
	if like.UserID != req.UserID {
		return errors.New("this like does not belong to this user")
	}
	if err := s.likesStorage.Delete(ctx, req.LikeID); err != nil {
		return err
	}
	return nil
}
