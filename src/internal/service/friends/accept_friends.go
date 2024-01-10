package friendsservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *friendService) AcceptFriends(ctx context.Context, friends model.Friends) error {
	fr, err := s.friendStorage.Get(ctx, friends.SenderID, friends.ReceiverID)
	if err != nil {
		return errors.New(constant.FriendRequestNotFound)
	}
	if fr.Accepted {
		return errors.New(constant.FriendAlreadyFriends)
	}
	if err := s.friendStorage.Accept(ctx, fr.SenderID, fr.ReceiverID); err != nil {
		return errors.New(fmt.Sprintf(constant.FriendAcceptFail, err.Error()))
	}
	return nil
}
