package postservice

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/repository/post"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type PostService interface {
	Create(ctx context.Context, req dto.CreatePostRequest) error
	Delete(ctx context.Context, req dto.DeletePostRequest) error
	Get(ctx context.Context, req dto.GetPostRequest) (dto.ResponseForGetPost, error)
	GetAll(ctx context.Context, req dto.GetAllPostRequest) ([]dto.ResponseForGetAllPost, error)
	Update(ctx context.Context, req dto.UpdatePostRequest) error
}

type postService struct {
	postStorage PostStorer
	repository  *routes.Repositories
	utils       *util.Util
}

type PostServiceOption func(*postService)

func WithPostServicePostStorage(postStorage PostStorer) PostServiceOption {
	return func(u *postService) {
		u.postStorage = postStorage
	}
}

func NewPostService(utils *util.Util, repo *routes.Repositories, opts ...PostServiceOption) PostService {
	u := &postService{
		utils:      utils,
		repository: repo,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
