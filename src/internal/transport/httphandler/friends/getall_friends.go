package friendshttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) GetAllFriends(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	if userID == 0 {
		return c.JSON(h.util.BasicError(constant.UnknowError, fiber.StatusUnauthorized))
	}
	res, err := h.service.GetAllFriends(c.Context(), userID)
	if err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, res))
}
