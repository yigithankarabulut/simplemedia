package posthttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
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
		return c.JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}
	if c.Locals("userID") == nil {
		return c.JSON(h.util.BasicError("You are not authorized to see this post", fiber.StatusUnauthorized))
	}
	userID := c.Locals("userID").(uint)
	req.UserID = userID
	res, err = h.service.Get(c.Context(), req)
	if err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, res))
}
