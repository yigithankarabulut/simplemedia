package likeshttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/likes/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	var (
		req dto.BaseLikesRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	if req.CommentID == 0 && req.PostID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, constant.PostOrCommentIDNotFound), fiber.StatusBadRequest))
	}
	req.UserID = c.Locals("userID").(uint)
	err := h.service.Create(c.Context(), req)
	if err != nil {
		return c.Status(404).JSON(h.util.Response(404, err.Error()))
	}
	return c.Status(200).JSON(h.util.Response(200, constant.AddLikeSuccess))
}
