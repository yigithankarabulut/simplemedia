package posthttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/apiserver/routes"
	. "github.com/yigithankarabulut/simplemedia/src/internal/service/post"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/basehttphandler"
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
	app.Get("/p", h.ShortRedirect)
	post := app.Group("/post")
	post.Use(basehttphandler.AuthMiddleware())
	post.Post("/create", h.Create)
	post.Put("/update", h.Update)
	post.Get("/get", h.Get)
	post.Delete("/delete", h.Delete)
	post.Get("/getall", h.GetAll)
}

func (h *Handler) ShortRedirect(c *fiber.Ctx) error {
	shortLink := c.Query("id")
	fmt.Println(shortLink)

	postID, userID, err := h.util.DecodeShortURL(shortLink)
	if err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}

	c.Locals("userID", userID)
	c.Locals("postID", postID)
	return h.Get(c)
}
