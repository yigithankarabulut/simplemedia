package commenthttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) GetAll(c *fiber.Ctx) error {
	var (
		req dto.GetAllCommentRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	if req.CommentID == 0 && req.PostID == 0 {
		return c.JSON(h.util.BasicError(constant.IDNotFound, fiber.StatusBadRequest))
	}
	if req.CommentID != 0 && req.PostID != 0 {
		return c.JSON(h.util.BasicError(constant.CommentNotUseTogether, fiber.StatusBadRequest))
	}
	req.UserID = c.Locals("userID").(uint)

	resp, err := h.service.GetAll(c.Context(), req)
	if err != nil {
		return c.JSON(h.util.BasicError(err.Error(), fiber.StatusBadRequest))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, resp))
}
