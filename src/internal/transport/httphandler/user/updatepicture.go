package userhttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) UpdateProfilePicture(c *fiber.Ctx) error {
	var (
		req dto.UpdatePictureRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.Status(400).JSON(h.util.BasicError(
			fmt.Sprintf(constant.Validate, err.Error()),
			fiber.StatusBadRequest))
	}
	req.UserID = c.Locals("userID").(uint)
	file, err := c.FormFile("picture")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(
			fmt.Sprintf(constant.Validate, err.Error()),
			fiber.StatusBadRequest))
	}
	filepath, err := h.util.SavePicture(file, "profile-picture", req.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	if err := c.SaveFile(file, filepath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	req.PictureUrl = filepath
	if err := h.service.UpdateProfilePicture(c.Context(), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.Status(fiber.StatusOK).JSON(h.util.Response(fiber.StatusOK, constant.SuccessUpdatePicture))
}
