package commenthttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
)

func (h *Handler) Get(c *fiber.Ctx) error {
	var (
		req dto.GetCommentRequest
		res dto.BaseCommentResponse
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}
	req.UserID = c.Locals("userID").(uint)

	res, err := h.service.Get(c.Context(), req)
	if err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, res))
}
