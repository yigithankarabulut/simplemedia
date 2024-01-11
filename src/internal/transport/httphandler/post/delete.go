package posthttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) Delete(c *fiber.Ctx) error {
	var (
		req dto.DeletePostRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	userID := c.Locals("userID").(uint)
	req.UserID = userID
	if err := h.service.Delete(c.Context(), req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.Response(fiber.StatusBadRequest, err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(h.util.Response(fiber.StatusOK, constant.SuccessDeletePost))
}
