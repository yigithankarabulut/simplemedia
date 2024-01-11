package friendshttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/friends/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
	"strconv"
)

func (h *Handler) AddFriends(c *fiber.Ctx) error {
	var (
		req dto.AddFriendRequest
	)
	strID := c.Query("receiver_id")
	if strID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.Response(fiber.StatusBadRequest, fmt.Sprintf(constant.Validate, constant.InvalidParam)))
	}
	id, err := strconv.Atoi(strID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.Response(fiber.StatusBadRequest, fmt.Sprintf(constant.Validate, err.Error())))
	}
	req.ReceiverID = uint(id)
	userID := c.Locals("userID").(uint)
	req.SenderID = userID
	if err := h.service.AddFriends(c.Context(), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.Status(fiber.StatusOK).JSON(h.util.Response(fiber.StatusOK, constant.FriendRequestSent))
}
