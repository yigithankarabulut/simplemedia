package userhttphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	var (
		req dto.LoginRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusNotFound))
	}
	resp, err := h.service.Login(c.Context(), req)
	if err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, resp))
}
