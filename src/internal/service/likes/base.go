package likesservice

import . "github.com/yigithankarabulut/simplemedia/src/internal/repository/likes"

type LikesService interface {
}

type likesService struct {
	likesStorage LikeStorer
}

type LikesServiceOption func(*likesService)

func WithLikesServiceLikesStorage(likesStorage LikeStorer) LikesServiceOption {
	return func(u *likesService) {
		u.likesStorage = likesStorage
	}
}

func NewLikesService(opts ...LikesServiceOption) LikesService {
	u := &likesService{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
