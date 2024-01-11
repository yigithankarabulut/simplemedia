package posthttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) Update(c *fiber.Ctx) error {
	var (
		req dto.UpdatePostRequest
		err error
	)
	if err = h.util.Validate(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	userID := c.Locals("userID").(uint)
	req.UserID = userID
	if err = h.service.Update(c.Context(), req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}
	return c.Status(fiber.StatusOK).JSON(h.util.Response(fiber.StatusOK, constant.SuccessUpdatePost))
}
