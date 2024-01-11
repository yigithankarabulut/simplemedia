package likesservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/likes/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *likesService) Create(ctx context.Context, req dto.BaseLikesRequest) error {
	var (
		like *model.Likes
	)
	like = req.Convert()
	if like.CommentID > 0 {
		if _, err := s.repositories.CommentStorer.GetByID(ctx, like.CommentID); err != nil {
			return errors.New(constant.CommentNotFound)
		}
		if _, err := s.likesStorage.GetComment(ctx, like.UserID, like.CommentID); err == nil {
			return errors.New(constant.CommentAlreadyLiked)
		}
	}
	if like.PostID > 0 {
		if _, err := s.repositories.PostStorer.GetByID(ctx, like.PostID); err != nil {
			return errors.New(constant.PostNotFound)
		}
		if _, err := s.likesStorage.GetPost(ctx, like.UserID, like.PostID); err == nil {
			return errors.New(constant.PostAlreadyLiked)
		}
	}
	if err := s.likesStorage.Insert(ctx, like); err != nil {
		return errors.New(fmt.Sprintf(constant.FailedLikeAdd, err.Error()))
	}
	return nil
}
