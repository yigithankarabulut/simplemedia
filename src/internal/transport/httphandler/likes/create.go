package likeshttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/likes/dto"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	var (
		req dto.BaseLikesRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(err.Error(), 404))
	}
	if req.CommentID == 0 && req.PostID == 0 {
		return c.JSON(h.util.BasicError("CommentID or PostID must be provided", 404))
	}
	req.UserID = c.Locals("userID").(uint)
	err := h.service.Create(c.Context(), req)
	if err != nil {
		return c.JSON(h.util.BasicError(err.Error(), 404))
	}
	return c.JSON(h.util.Response(200, "Like created successfully"))
}
