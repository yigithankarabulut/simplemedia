package usersevice

import . "github.com/yigithankarabulut/simplemedia/src/internal/repository/user"

type UserService interface {
}

type userService struct {
	userStorage UserStorer
}

type UserServiceOption func(*userService)

func WithUserServiceUserStorage(userStorage UserStorer) UserServiceOption {
	return func(u *userService) {
		u.userStorage = userStorage
	}
}

func NewUserService(opts ...UserServiceOption) UserService {
	u := &userService{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
