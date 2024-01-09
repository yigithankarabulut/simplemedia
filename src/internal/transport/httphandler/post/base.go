package posthttphandler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yigithankarabulut/simplemedia/src/apiserver"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/post"
)

type Handler struct {
	service PostService
}

func NewHandler(service PostService) apiserver.Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes(app fiber.Router) {

}
