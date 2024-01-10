package commentservice

import (
	. "github.com/yigithankarabulut/simplemedia/src/internal/repository/comment"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type CommentService interface {
}

type commentService struct {
	commentStorage CommentStorer
	util           *util.Util
}

type CommentServiceOption func(*commentService)

func WithCommentServiceCommentStorage(commentStorage CommentStorer) CommentServiceOption {
	return func(u *commentService) {
		u.commentStorage = commentStorage
	}
}

func NewCommentService(util *util.Util, opts ...CommentServiceOption) CommentService {
	u := &commentService{
		util: util,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
