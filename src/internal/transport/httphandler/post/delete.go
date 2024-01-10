package posthttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
)

func (h *Handler) Delete(c *fiber.Ctx) error {
	var (
		req dto.DeletePostRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}
	userID := c.Locals("userID").(uint)
	req.UserID = userID
	if err := h.service.Delete(c.Context(), req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, "Post deleted successfully"))
}
