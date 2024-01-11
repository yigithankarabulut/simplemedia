package commentservice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
)

func (s *commentService) Update(ctx context.Context, req dto.UpdateCommentRequest) error {
	var (
		comment *model.Comment
		err     error
	)
	comment, err = s.commentStorage.GetByID(ctx, req.ID)
	if err != nil {
		return errors.New("comment not found")
	}
	if comment.UserID != req.UserID {
		return errors.New("you are not owner of this comment")
	}
	comment.Content = req.Content
	if err = s.commentStorage.Update(ctx, comment); err != nil {
		return err
	}
	return nil
}
