package commenthttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
)

func (h *Handler) Update(c *fiber.Ctx) error {
	var (
		req dto.UpdateCommentRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}
	req.UserID = c.Locals("userID").(uint)

	if err := h.service.Update(c.Context(), req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, "Comment updated successfully"))
}
