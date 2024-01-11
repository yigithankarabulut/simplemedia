package friendshttphandler

import "github.com/gofiber/fiber/v2"

func (h *Handler) GetAllFriends(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	res, err := h.service.GetAllFriends(c.Context(), userID)
	if err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, res))
}
