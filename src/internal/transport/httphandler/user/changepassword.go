package userhttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
)

func (h *Handler) ChangePassword(c *fiber.Ctx) error {
	var (
		req dto.ChangePwdRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusNotFound))
	}
	userID := c.Locals("userID").(uint)
	if userID == 0 {
		return c.JSON(h.util.BasicError("User not found", fiber.StatusNotFound))
	}
	if err := h.service.ChangePassword(c.Context(), userID, req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, "Password changed successfully"))
}
