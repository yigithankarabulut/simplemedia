package usersevice

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/repository/user"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest, errCh chan error) error
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	Logout(ctx context.Context, userID uint) error
	ChangePassword(ctx context.Context, userID uint, req dto.ChangePwdRequest) error
	UpdateProfilePicture(ctx context.Context, req dto.UpdatePictureRequest) error
}

type userService struct {
	userStorage UserStorer
	repository  *routes.Repositories
	utils       *util.Util
}

type UserServiceOption func(*userService)

func WithUserServiceUserStorage(userStorage UserStorer) UserServiceOption {
	return func(u *userService) {
		u.userStorage = userStorage
	}
}

func NewUserService(util *util.Util, repo *routes.Repositories, opts ...UserServiceOption) UserService {
	u := &userService{
		utils:      util,
		repository: repo,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
