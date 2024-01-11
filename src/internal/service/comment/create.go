package commentservice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
)

func (s *commentService) Create(ctx context.Context, req dto.CreateCommentRequest) error {
	var (
		comment *model.Comment
		err     error
	)
	post, err := s.repository.PostStorer.GetByID(ctx, req.PostID)
	if err != nil {
		return errors.New("post not found")
	}
	if post.UserID != req.UserID {

		fr, err := s.repository.FriendStorer.Get(ctx, req.UserID, post.UserID)
		if err != nil {
			fr2, err := s.repository.FriendStorer.Get(ctx, post.UserID, req.UserID)
			if err != nil {
				return errors.New("you are not friends")
			}
			fr = fr2
		}
		if !fr.Accepted {
			return errors.New("you are not friends")
		}
	}
	if req.ParentID > 0 {
		_, err = s.commentStorage.GetByID(ctx, req.ParentID)
		if err != nil {
			return errors.New("parent comment not found")
		}
	}
	comment = req.Convert()
	if err = s.commentStorage.Insert(ctx, comment); err != nil {
		return err
	}
	return nil
}
