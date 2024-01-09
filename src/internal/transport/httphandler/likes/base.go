package likeshttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/likes"
)

type Handler struct {
	service LikesService
}

func NewHandler(service LikesService) routes.Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes(app fiber.Router) {
	likes := app.Group("/likes")
	//app.Use()

	likes.Post("/create", h.Create)
	likes.Delete("/delete", h.Delete)
	likes.Get("/get", h.Get)
	likes.Get("/getall", h.GetAll)
}
