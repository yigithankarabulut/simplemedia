package commentservice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
)

func (s *commentService) Get(ctx context.Context, req dto.GetCommentRequest) (dto.BaseCommentResponse, error) {
	var (
		res     dto.BaseCommentResponse
		comment *model.Comment
		err     error
	)
	comment, err = s.commentStorage.GetByID(ctx, req.ID)
	if err != nil {
		return dto.BaseCommentResponse{}, errors.New("comment not found")
	}
	fr, err := s.repository.FriendStorer.Get(ctx, req.UserID, comment.UserID)
	if err != nil {
		fr2, err := s.repository.FriendStorer.Get(ctx, comment.UserID, req.UserID)
		if err != nil {
			return dto.BaseCommentResponse{}, errors.New("you are not friends")
		}
		fr = fr2
	}
	if !fr.Accepted {
		return dto.BaseCommentResponse{}, errors.New("you are not friends")
	}
	res.Convert(comment)
	return res, nil
}
