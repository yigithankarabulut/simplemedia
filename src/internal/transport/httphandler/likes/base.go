package likeshttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/likes"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type Handler struct {
	service LikesService
	util    *util.Util
}

func NewHandler(util *util.Util, service LikesService) routes.Handler {
	return &Handler{
		service: service,
		util:    util,
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
