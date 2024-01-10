package postservice

import (
	. "github.com/yigithankarabulut/simplemedia/src/internal/repository/post"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type PostService interface {
}

type postService struct {
	postStorage PostStorer
	utils       *util.Util
}

type PostServiceOption func(*postService)

func WithPostServicePostStorage(postStorage PostStorer) PostServiceOption {
	return func(u *postService) {
		u.postStorage = postStorage
	}
}

func NewPostService(utils *util.Util, opts ...PostServiceOption) PostService {
	u := &postService{
		utils: utils,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
