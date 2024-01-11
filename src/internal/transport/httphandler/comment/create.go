package commenthttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	var (
		req dto.CreateCommentRequest
		err error
	)
	if err = h.util.Validate(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	req.UserID = c.Locals("userID").(uint)
	file, err := c.FormFile("image")
	if err == nil {
		filepath, err := h.util.SavePicture(file, req.Content, "comment-images") // TODO:
		if err != nil {
			return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
		}
		if err := c.SaveFile(file, filepath); err != nil {
			return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
		}
		req.ImageUrl = filepath
	}
	err = h.service.Create(c.Context(), req)
	if err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, constant.CreateCommentSuccess))
}
