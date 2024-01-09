package friendshttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/friends"
)

type Handler struct {
	service FriendService
}

func NewHandler(service FriendService) routes.Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes(app fiber.Router) {
	friends := app.Group("/friends")
	//app.Use()

	friends.Post("/add", h.AddFriends)
	friends.Post("/accept", h.AcceptFriends)
	friends.Post("/decline", h.DeclineFriends)
	friends.Delete("/delete", h.DeleteFriends)
	friends.Get("/request", h.GetAllFriendRequests)
}
