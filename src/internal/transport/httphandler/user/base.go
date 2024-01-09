package userhttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/user"
)

type Handler struct {
	service UserService
}

func NewHandler(service UserService) routes.Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes(app fiber.Router) {
	app.Post("/register", h.Register)
	app.Post("/login", h.Login)
	user := app.Group("/user")
	//user.Use()
	user.Post("/logout", h.Logout)
	user.Put("/pwd", h.ChangePassword)
	user.Put("/picture", h.UpdateProfilePicture)
}
