package postservice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *postService) Get(ctx context.Context, req dto.GetPostRequest) (dto.ResponseForGetPost, error) {
	var (
		res dto.ResponseForGetPost
		err error
	)
	post, err := s.postStorage.GetByID(ctx, req.ID)
	if err != nil {
		return res, err
	}
	if post.UserID != req.UserID {
		return res, errors.New(constant.FailedGetPost)
	}
	res.ID = post.ID
	res.UserID = post.UserID
	res.Title = post.Title
	res.Content = post.Content
	res.ImageUrl = post.Image
	res.ShortUrl = s.utils.EncodeToShortURL(post.ID, post.UserID)

	return res, nil
}
