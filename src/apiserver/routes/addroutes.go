package routes

import (
	"github.com/gofiber/fiber/v2"
	commentstorage "github.com/yigithankarabulut/simplemedia/src/internal/repository/comment"
	friendsstorage "github.com/yigithankarabulut/simplemedia/src/internal/repository/friends"
	likesstorage "github.com/yigithankarabulut/simplemedia/src/internal/repository/likes"
	poststorage "github.com/yigithankarabulut/simplemedia/src/internal/repository/post"
	userstorage "github.com/yigithankarabulut/simplemedia/src/internal/repository/user"
)

type Handler interface {
	AddRoutes(router fiber.Router)
}

type Repositories struct {
	commentstorage.CommentStorer
	poststorage.PostStorer
	likesstorage.LikeStorer
	friendsstorage.FriendStorer
	userstorage.UserStorer
}

func NewRepositories() *Repositories {
	return &Repositories{}
}
