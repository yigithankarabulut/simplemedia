package commentservice

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
)

func (s *commentService) GetAll(ctx context.Context, req dto.GetAllCommentRequest) ([]dto.BaseCommentResponse, error) {
	var (
		res        []dto.BaseCommentResponse
		resComment []model.Comment
		err        error
	)
	if req.CommentID != 0 {
		resComment, err = s.repository.CommentStorer.GetAllByCommentID(ctx, req.CommentID)
	}
	if req.PostID != 0 {
		resComment, err = s.repository.CommentStorer.GetAllByPostID(ctx, req.PostID)
	}
	if err != nil {
		return nil, err
	}
	for _, comment := range resComment {
		var (
			user *model.User
		)
		user, err = s.repository.UserStorer.GetByID(ctx, comment.UserID)
		if err != nil {
			return nil, err
		}
		res = append(res, dto.BaseCommentResponse{
			ID:       comment.ID,
			Username: user.Username,
			Image:    comment.Image,
			ParentID: comment.ParentID,
			Content:  comment.Content,
		})
	}
	return res, nil
}
