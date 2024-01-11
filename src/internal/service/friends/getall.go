package friendsservice

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/friends/dto"
)

func (s *friendService) GetAllFriendRequests(ctx context.Context, userID uint) ([]dto.GetAllFriendResponse, error) {
	var (
		res       []dto.GetAllFriendResponse
		friendReq []model.Friends
		err       error
	)

	friendReq, err = s.friendStorage.GetAllRequests(ctx, userID)
	if err != nil {
		return nil, err
	}
	for _, friend := range friendReq {
		var (
			friendID uint
			user     *model.User
		)
		if friend.SenderID == userID {
			friendID = friend.ReceiverID
		} else {
			friendID = friend.SenderID
		}
		user, err := s.repository.UserStorer.GetByID(ctx, friendID)
		if err != nil {
			return nil, err
		}
		res = append(res, dto.GetAllFriendResponse{
			ID:       friendID,
			Username: user.Username,
		})
	}
	return res, nil
}

func (s *friendService) GetAllFriends(ctx context.Context, userID uint) ([]dto.GetAllFriendResponse, error) {
	var (
		res []dto.GetAllFriendResponse
		err error
	)
	resp, err := s.friendStorage.GetAllFriends(ctx, userID)
	if err != nil {
		return nil, err
	}
	for _, friend := range resp {
		var (
			friendID uint
			user     *model.User
		)
		if friend.SenderID == userID {
			friendID = friend.ReceiverID
		} else {
			friendID = friend.SenderID
		}
		user, err = s.repository.UserStorer.GetByID(ctx, friendID)
		if err != nil {
			return nil, err
		}
		res = append(res, dto.GetAllFriendResponse{
			ID:       friendID,
			Username: user.Username,
		})
	}
	return res, nil
}
