package commenthttphandler

import (
	"github.com/gofiber/fiber/v3"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/comment"
)

type Handler struct {
	service CommentService
}

func NewHandler(service CommentService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes(router fiber.Router) {

}
