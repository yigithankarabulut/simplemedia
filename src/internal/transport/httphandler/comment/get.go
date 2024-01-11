package commenthttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) Get(c *fiber.Ctx) error {
	var (
		req dto.GetCommentRequest
		res dto.BaseCommentResponse
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	req.UserID = c.Locals("userID").(uint)

	res, err := h.service.Get(c.Context(), req)
	if err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, res))
}
