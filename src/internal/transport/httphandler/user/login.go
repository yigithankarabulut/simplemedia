package userhttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK).JSON(util.Utils.Response(201, "okkeyys"))
	return nil
}
