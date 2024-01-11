package commenthttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
)

func (h *Handler) GetAll(c *fiber.Ctx) error {
	var (
		req dto.GetAllCommentRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(err.Error(), fiber.StatusBadRequest))
	}
	if req.CommentID == 0 && req.PostID == 0 {
		return c.JSON(h.util.BasicError("CommentID or PostID is required", fiber.StatusBadRequest))
	}
	if req.CommentID != 0 && req.PostID != 0 {
		return c.JSON(h.util.BasicError("CommentID and PostID cannot be used together", fiber.StatusBadRequest))
	}
	req.UserID = c.Locals("userID").(uint)

	resp, err := h.service.GetAll(c.Context(), req)
	if err != nil {
		return c.JSON(h.util.BasicError(err.Error(), fiber.StatusBadRequest))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, resp))
}
