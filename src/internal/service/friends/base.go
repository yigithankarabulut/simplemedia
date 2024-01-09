package friendsservice

import . "github.com/yigithankarabulut/simplemedia/src/internal/repository/friends"

type FriendService interface {
}

type friendService struct {
	friendStorage FriendStorer
}

type FriendServiceOption func(*friendService)

func WithFriendServiceFriendStorage(friendStorage FriendStorer) FriendServiceOption {
	return func(u *friendService) {
		u.friendStorage = friendStorage
	}
}

func NewFriendService(opts ...FriendServiceOption) FriendService {
	u := &friendService{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
