package posthttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	var (
		req dto.CreatePostRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	userID := c.Locals("userID").(uint)
	req.UserID = userID
	file, err := c.FormFile("image")
	if err == nil {
		filepath, err := h.util.SavePicture(file, req.Title, "post-images") // TODO: change title to post id
		if err != nil {
			return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
		}
		if err := c.SaveFile(file, filepath); err != nil {
			return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
		}
		req.ImageUrl = filepath
	}
	if err := h.service.Create(c.Context(), req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, "Post created successfully"))
}
