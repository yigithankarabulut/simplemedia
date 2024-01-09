package commenthttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/comment"
)

type Handler struct {
	service CommentService
}

func NewHandler(service CommentService) routes.Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes(app fiber.Router) {
	comment := app.Group("/comment")
	//comment.Use()
	comment.Post("/create", h.Create)
	comment.Put("/update", h.Update)
	comment.Delete("/delete", h.Delete)
	comment.Get("/get", h.Get)
	comment.Get("/getall", h.GetAll)
}
