package friendshttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/friends"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/basehttphandler"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type Handler struct {
	service FriendService
	util    *util.Util
}

func NewHandler(util *util.Util, service FriendService) routes.Handler {
	return &Handler{
		service: service,
		util:    util,
	}
}

func (h *Handler) AddRoutes(app fiber.Router) {
	friends := app.Group("/friends")
	app.Use(basehttphandler.AuthMiddleware())
	friends.Post("/add", h.AddFriends)
	friends.Post("/accept", h.AcceptFriends)
	friends.Delete("/decline", h.DeclineFriends)
	friends.Delete("/delete", h.DeleteFriends)
	friends.Get("/request", h.GetAllFriendRequests)
}
