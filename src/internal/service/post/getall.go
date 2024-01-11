package postservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *postService) GetAll(ctx context.Context, req dto.GetAllPostRequest) ([]dto.ResponseForGetAllPost, error) {
	var (
		res []dto.ResponseForGetAllPost
		err error
	)
	if _, err := s.repository.UserStorer.GetByID(ctx, req.ID); err != nil {
		return res, errors.New("user not found")
	}
	if req.UserID != req.ID {
		fr, err := s.repository.FriendStorer.Get(ctx, req.ID, req.UserID)
		if err != nil {
			fr2, err := s.repository.FriendStorer.Get(ctx, req.UserID, req.ID)
			if err != nil {
				return res, errors.New(constant.FailedFriendPost)
			}
			fr = fr2
		}
		if !fr.Accepted {
			return res, errors.New(constant.FailedFriendPost)
		}
	}
	resp, err := s.postStorage.GetByUserID(ctx, req.ID)
	if err != nil {
		return res, errors.New(fmt.Sprintf(constant.FailedPostGet, err.Error()))
	}
	for _, v := range resp {
		res = append(res, dto.ResponseForGetAllPost{
			ID:       v.ID,
			UserID:   v.UserID,
			Content:  v.Content,
			Title:    v.Title,
			ImageUrl: v.Image,
		})
	}
	return res, nil
}
