package postservice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
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
		return res, errors.New("You are not authorized to see this post")
	}
	res.ID = post.ID
	res.UserID = post.UserID
	res.Title = post.Title
	res.Content = post.Content
	res.ImageUrl = post.Image

	res.ShortUrl = s.utils.CreateShortLink(res.ID, res.UserID)

	return res, nil
}
