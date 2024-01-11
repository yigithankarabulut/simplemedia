package commentservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *commentService) Delete(ctx context.Context, req dto.DeleteCommentRequest) error {
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
	if err = s.commentStorage.Delete(ctx, req.ID); err != nil {
		return errors.New(fmt.Sprintf(constant.CommentDeleteFail, err.Error()))
	}
	return nil
}
