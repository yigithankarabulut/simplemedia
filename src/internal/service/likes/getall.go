package likesservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/likes/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *likesService) GetAll(ctx context.Context, req dto.BaseLikesRequest) (dto.GetAllLikesResponse, error) {
	var (
		likes        []model.Likes
		likeComments []model.Likes
		likePosts    []model.Likes
		baseResponse []dto.BaseLikesResponse
		res          dto.GetAllLikesResponse
		err          error
		count        uint
	)

	if req.CommentID != 0 {
		likeComments, err = s.likesStorage.GetAllComment(ctx, req.CommentID)
		if err != nil {
			return dto.GetAllLikesResponse{}, errors.New(fmt.Sprintf(constant.FailedLikeList, err.Error()))
		}
	}
	if req.PostID != 0 {
		likePosts, err = s.likesStorage.GetAllPost(ctx, req.PostID)
		if err != nil {
			return dto.GetAllLikesResponse{}, errors.New(fmt.Sprintf(constant.FailedLikeList, err.Error()))
		}
	}
	likes = append(likeComments, likePosts...)
	count = uint(len(likes))
	for _, like := range likes {
		baseResponse = append(baseResponse, dto.BaseLikesResponse{
			ID:     like.ID,
			UserID: like.UserID,
		})
	}
	res = dto.GetAllLikesResponse{
		Count: count,
		Likes: baseResponse,
	}
	return res, nil
}
