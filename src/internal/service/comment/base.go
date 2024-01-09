package commentservice

import . "github.com/yigithankarabulut/simplemedia/src/internal/repository/comment"

type CommentService interface {
}

type commentService struct {
	commentStorage CommentStorer
}

type CommentServiceOption func(*commentService)

func WithCommentServiceCommentStorage(commentStorage CommentStorer) CommentServiceOption {
	return func(u *commentService) {
		u.commentStorage = commentStorage
	}
}

func NewCommentService(opts ...CommentServiceOption) CommentService {
	u := &commentService{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
