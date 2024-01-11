package commentservice

import (
	"context"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/repository/comment"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type CommentService interface {
	Create(ctx context.Context, req dto.CreateCommentRequest) error
	Update(ctx context.Context, req dto.UpdateCommentRequest) error
	Delete(ctx context.Context, req dto.DeleteCommentRequest) error
	Get(ctx context.Context, req dto.GetCommentRequest) (dto.BaseCommentResponse, error)
	GetAll(ctx context.Context, req dto.GetAllCommentRequest) ([]dto.BaseCommentResponse, error)
}

type commentService struct {
	commentStorage CommentStorer
	repository     *routes.Repositories
	util           *util.Util
}

type CommentServiceOption func(*commentService)

func WithCommentServiceCommentStorage(commentStorage CommentStorer) CommentServiceOption {
	return func(u *commentService) {
		u.commentStorage = commentStorage
	}
}

func NewCommentService(util *util.Util, repo *routes.Repositories, opts ...CommentServiceOption) CommentService {
	u := &commentService{
		util:       util,
		repository: repo,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
