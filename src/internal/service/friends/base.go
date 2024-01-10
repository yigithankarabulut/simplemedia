package friendsservice

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	. "github.com/yigithankarabulut/simplemedia/src/internal/repository/friends"
	. "github.com/yigithankarabulut/simplemedia/src/internal/repository/user"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/friends/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type FriendService interface {
	AddFriends(ctx context.Context, req dto.AddFriendRequest) error
	AcceptFriends(ctx context.Context, fr model.Friends) error
	DeclineFriends(ctx context.Context, fr model.Friends) error
	DeleteFriends(ctx context.Context, fr model.Friends) error
}

type friendService struct {
	friendStorage FriendStorer
	userStorage   UserStorer
	utils         *util.Util
}

type FriendServiceOption func(*friendService)

func WithFriendServiceFriendStorage(friendStorage FriendStorer) FriendServiceOption {
	return func(u *friendService) {
		u.friendStorage = friendStorage
	}
}

func NewFriendService(util *util.Util, userStorage UserStorer, opts ...FriendServiceOption) FriendService {
	u := &friendService{
		utils:       util,
		userStorage: userStorage,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
