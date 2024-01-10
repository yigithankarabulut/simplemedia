package friendsservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *friendService) DeleteFriends(ctx context.Context, friends model.Friends) error {
	fr, err := s.friendStorage.Get(ctx, friends.SenderID, friends.ReceiverID)
	fr2, err2 := s.friendStorage.Get(ctx, friends.ReceiverID, friends.SenderID)
	if err != nil && err2 != nil {
		return errors.New(constant.FriendNotFriends)
	} else if err != nil {
		fr = fr2
	}
	if !fr.Accepted {
		return errors.New(constant.FriendNotFriends + " Please use decline friends")
	}
	if err := s.friendStorage.Delete(ctx, fr.SenderID, fr.ReceiverID); err != nil {
		return errors.New(fmt.Sprintf(constant.FriendDeleteFail, err.Error()))
	}
	return nil
}
