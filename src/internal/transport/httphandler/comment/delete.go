package commenthttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
)

func (h *Handler) Delete(c *fiber.Ctx) error {
	var (
		req dto.DeleteCommentRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}
	req.UserID = c.Locals("userID").(uint)

	if err := h.service.Delete(c.Context(), req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, "Comment deleted successfully"))
}
