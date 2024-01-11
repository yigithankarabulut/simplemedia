package posthttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) Get(c *fiber.Ctx) error {
	var (
		req dto.GetPostRequest
		res dto.ResponseForGetPost
		err error
	)
	if c.Locals("postID") != nil {
		req.ID = c.Locals("postID").(uint)
	} else if err = h.util.Validate(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	if c.Locals("userID") == nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(constant.UnknowError, fiber.StatusBadRequest))
	}
	userID := c.Locals("userID").(uint)
	req.UserID = userID
	res, err = h.service.Get(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}
	return c.Status(fiber.StatusOK).JSON(h.util.Response(fiber.StatusOK, res))
}
