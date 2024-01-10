package userhttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (h *Handler) Register(c *fiber.Ctx) error {
	var (
		req dto.RegisterRequest
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusNotFound))
	}
	if err := h.service.Register(c.Context(), req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, "User registered successfully"))
}
