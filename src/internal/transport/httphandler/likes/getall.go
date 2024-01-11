package likeshttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/likes/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) GetAll(c *fiber.Ctx) error {
	var (
		req dto.BaseLikesRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	if req.CommentID == 0 && req.PostID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(constant.IDNotFound, fiber.StatusBadRequest))
	}
	req.UserID = c.Locals("userID").(uint)
	likes, err := h.service.GetAll(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.Response(fiber.StatusBadRequest, err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(h.util.Response(fiber.StatusOK, likes))
}
