package userhttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	var (
		req dto.LoginRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.Status(400).JSON(h.util.BasicError(
			fmt.Sprintf(constant.Validate, err.Error()),
			fiber.StatusBadRequest))
	}
	resp, err := h.service.Login(c.Context(), req)
	if err != nil {
		return c.Status(400).JSON(h.util.BasicError(err, fiber.StatusBadRequest))
	}
	return c.Status(200).JSON(h.util.Response(fiber.StatusOK, resp))
}
