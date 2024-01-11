package friendshttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
	"strconv"
)

func (h *Handler) AcceptFriends(c *fiber.Ctx) error {
	var (
		friend model.Friends
	)
	strID := c.Query("id")
	if strID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, constant.InvalidParam), fiber.StatusBadRequest))
	}
	id, err := strconv.Atoi(strID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.Response(fiber.StatusBadRequest, fmt.Sprintf(constant.Validate, err.Error())))
	}
	userID := c.Locals("userID").(uint)
	if userID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(h.util.Response(fiber.StatusBadRequest, fmt.Sprintf(constant.Validate, "unauthorized")))
	}
	friend.SenderID = uint(id)
	friend.ReceiverID = userID
	if err := h.service.AcceptFriends(c.Context(), friend); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.Status(fiber.StatusOK).JSON(h.util.Response(fiber.StatusOK, constant.FriendRequestAccept))
}
