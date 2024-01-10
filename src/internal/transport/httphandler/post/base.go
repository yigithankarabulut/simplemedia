package posthttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/post"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

type Handler struct {
	service PostService
	util    *util.Util
}

func NewHandler(util *util.Util, service PostService) routes.Handler {
	return &Handler{
		service: service,
		util:    util,
	}
}

func (h *Handler) AddRoutes(app fiber.Router) {
	post := app.Group("/post")
	//post.Use()
	post.Post("/create", h.Create) // short url generate
	post.Put("/update", h.Update)  // ilk 5 dk içerisinde
	post.Delete("/delete", h.Delete)
	post.Get("/get", h.Get)
	post.Get("/getall", h.GetAll)
}
