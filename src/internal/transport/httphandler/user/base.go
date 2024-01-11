package userhttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/user"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/basehttphandler"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type Handler struct {
	service UserService
	util    *util.Util
}

func NewHandler(util *util.Util, service UserService) routes.Handler {
	return &Handler{
		service: service,
		util:    util,
	}
}

func (h *Handler) AddRoutes(app fiber.Router) {
	app.Post("/register", h.Register)
	app.Post("/login", h.Login)
	user := app.Group("/user")
	user.Use(basehttphandler.AuthMiddleware())
	user.Post("/logout", h.Logout)
	user.Put("/pwd", h.ChangePassword)
	user.Put("/picture", h.UpdateProfilePicture)
}
