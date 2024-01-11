package userhttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
)

func (h *Handler) UpdateProfilePicture(c *fiber.Ctx) error {
	var (
		req dto.UpdatePictureRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}
	req.UserID = c.Locals("userID").(uint)
	file, err := c.FormFile("picture")
	if err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}
	filepath, err := h.util.SavePicture(file, "profile-picture", req.Username)
	if err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	if err := c.SaveFile(file, filepath); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	req.PictureUrl = filepath
	if err := h.service.UpdateProfilePicture(c.Context(), req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, "Profile picture updated successfully"))
}
