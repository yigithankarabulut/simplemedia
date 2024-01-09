package postservice

import . "github.com/yigithankarabulut/simplemedia/src/internal/repository/post"

type PostService interface {
}

type postService struct {
	postStorage PostStorer
}

type PostServiceOption func(*postService)

func WithPostServicePostStorage(postStorage PostStorer) PostServiceOption {
	return func(u *postService) {
		u.postStorage = postStorage
	}
}

func NewPostService(opts ...PostServiceOption) PostService {
	u := &postService{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
