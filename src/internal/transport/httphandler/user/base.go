package userhttphandler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yigithankarabulut/simplemedia/src/apiserver"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/user"
)

type Handler struct {
	service UserService
}

func NewHandler(service UserService) apiserver.Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes(app fiber.Router) {
	app.Post("/register", h.Register)
	app.Post("/login", h.Login)
	user := app.Group("/user")
	user.Use()
	user.Post("/logout", h.Logout)
	user.Put("/pass", h.RenewPassword)
	user.Put("/picture", h.UpdateProfilePicture)
}

func (h *Handler) Register(c fiber.Ctx) error {
	return nil
}

func (h *Handler) Login(c fiber.Ctx) error {
	return nil
}

func (h *Handler) Logout(c fiber.Ctx) error {
	return nil
}

func (h *Handler) RenewPassword(c fiber.Ctx) error {
	return nil
}

func (h *Handler) UpdateProfilePicture(c fiber.Ctx) error {
	return nil
}
