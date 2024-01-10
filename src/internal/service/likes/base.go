package likesservice

import (
	. "github.com/yigithankarabulut/simplemedia/src/internal/repository/likes"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type LikesService interface {
}

type likesService struct {
	likesStorage LikeStorer
	util         *util.Util
}

type LikesServiceOption func(*likesService)

func WithLikesServiceLikesStorage(likesStorage LikeStorer) LikesServiceOption {
	return func(u *likesService) {
		u.likesStorage = likesStorage
	}
}

func NewLikesService(util *util.Util, opts ...LikesServiceOption) LikesService {
	u := &likesService{
		util: util,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
