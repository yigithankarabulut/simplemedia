package friendsservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/friends/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *friendService) AddFriends(ctx context.Context, req dto.AddFriendRequest) error {
	var (
		friend model.Friends
	)
	if req.SenderID == req.ReceiverID {
		return errors.New(constant.FriendCannotAddSelf)
	}
	friend.SenderID = req.SenderID
	friend.ReceiverID = req.ReceiverID
	if _, err := s.repository.UserStorer.GetByID(ctx, friend.ReceiverID); err != nil {
		return errors.New(constant.FriendUserNotFound)
	}
	res, err := s.friendStorage.Get(ctx, req.SenderID, req.ReceiverID)
	if err == nil {
		return errors.New(constant.FriendAlreadyExists)
	}

	if res.Accepted {
		return errors.New(constant.FriendAlreadyFriends)
	}

	friend.Accepted = false
	if err := s.friendStorage.Insert(ctx, &friend); err != nil {
		return errors.New(fmt.Sprintf(constant.FriendRequestSentFail, err.Error()))
	}
	return nil
}
