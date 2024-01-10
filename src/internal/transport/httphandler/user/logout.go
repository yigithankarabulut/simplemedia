package userhttphandler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Logout(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	if userID == 0 {
		return c.JSON(h.util.BasicError("User not found", fiber.StatusNotFound))
	}
	if err := h.service.Logout(c.Context(), userID); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, "User logged out successfully"))
}
