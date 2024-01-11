package userhttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) Logout(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	if userID == 0 {
		return c.Status(401).JSON(h.util.BasicError(
			constant.Unauthorized, fiber.StatusUnauthorized))
	}
	if err := h.service.Logout(c.Context(), userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.util.BasicError(err.Error(), fiber.StatusInternalServerError))
	}
	return c.Status(fiber.StatusOK).JSON(h.util.Response(fiber.StatusOK, "Logout successfully"))
}
