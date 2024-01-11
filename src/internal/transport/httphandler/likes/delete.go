package likeshttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/likes/dto"
)

func (h *Handler) Delete(c *fiber.Ctx) error {
	var (
		req dto.DeleteLikesRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(err.Error(), 404))
	}
	req.UserID = c.Locals("userID").(uint)

	err := h.service.Delete(c.Context(), req)
	if err != nil {
		return c.JSON(h.util.BasicError(err.Error(), 404))
	}
	return c.JSON(h.util.Response(200, "Like deleted successfully"))
}
