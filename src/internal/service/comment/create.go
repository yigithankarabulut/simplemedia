package commentservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *commentService) Create(ctx context.Context, req dto.CreateCommentRequest) error {
	var (
		comment *model.Comment
		err     error
	)
	post, err := s.repository.PostStorer.GetByID(ctx, req.PostID)
	if err != nil {
		return errors.New(constant.PostNotFound)
	}
	if post.UserID != req.UserID {

		fr, err := s.repository.FriendStorer.Get(ctx, req.UserID, post.UserID)
		if err != nil {
			fr2, err := s.repository.FriendStorer.Get(ctx, post.UserID, req.UserID)
			if err != nil {
				return errors.New(constant.FriendNotFound)
			}
			fr = fr2
		}
		if !fr.Accepted {
			return errors.New(constant.FriendNotFound)
		}
	}
	if req.ParentID > 0 {
		_, err = s.commentStorage.GetByID(ctx, req.ParentID)
		if err != nil {
			return errors.New(constant.CommentParentNotFound)
		}
	}
	comment = req.Convert()
	if err = s.commentStorage.Insert(ctx, comment); err != nil {
		return errors.New(fmt.Sprintf(constant.CommentCreateFail, err.Error()))
	}
	return nil
}
