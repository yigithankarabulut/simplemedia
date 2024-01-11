package likesservice

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/repository/likes"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/likes/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type LikesService interface {
	Create(ctx context.Context, req dto.BaseLikesRequest) error
	Delete(ctx context.Context, req dto.DeleteLikesRequest) error
	GetAll(ctx context.Context, req dto.BaseLikesRequest) (dto.GetAllLikesResponse, error)
}

type likesService struct {
	likesStorage LikeStorer
	repositories *routes.Repositories
	util         *util.Util
}

type LikesServiceOption func(*likesService)

func WithLikesServiceLikesStorage(likesStorage LikeStorer) LikesServiceOption {
	return func(u *likesService) {
		u.likesStorage = likesStorage
	}
}

func NewLikesService(util *util.Util, repo *routes.Repositories, opts ...LikesServiceOption) LikesService {
	u := &likesService{
		util:         util,
		repositories: repo,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
