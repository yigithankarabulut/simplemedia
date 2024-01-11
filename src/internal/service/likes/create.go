package likesservice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/likes/dto"
)

func (s *likesService) Create(ctx context.Context, req dto.BaseLikesRequest) error {
	var (
		like *model.Likes
	)
	like = req.Convert()
	if like.CommentID > 0 {
		if _, err := s.repositories.CommentStorer.GetByID(ctx, like.CommentID); err != nil {
			return errors.New("comment not found")
		}
		if _, err := s.likesStorage.GetComment(ctx, like.UserID, like.CommentID); err == nil {
			return errors.New("comment already liked")
		}
	}
	if like.PostID > 0 {
		if _, err := s.repositories.PostStorer.GetByID(ctx, like.PostID); err != nil {
			return errors.New("post not found")
		}
		if _, err := s.likesStorage.GetPost(ctx, like.UserID, like.PostID); err == nil {
			return errors.New("post already liked")
		}
	}
	if err := s.likesStorage.Insert(ctx, like); err != nil {
		return err
	}
	return nil
}
