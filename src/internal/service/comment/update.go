package commentservice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *commentService) Update(ctx context.Context, req dto.UpdateCommentRequest) error {
	var (
		comment *model.Comment
		err     error
	)
	comment, err = s.commentStorage.GetByID(ctx, req.ID)
	if err != nil {
		return errors.New(constant.CommentNotFound)
	}
	if comment.UserID != req.UserID {
		return errors.New(constant.CommentNotOwner)
	}
	comment.Content = req.Content
	if err = s.commentStorage.Update(ctx, comment); err != nil {
		return errors.New(constant.CommentUpdateFail)
	}
	return nil
}
