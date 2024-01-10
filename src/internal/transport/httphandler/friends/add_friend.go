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
		return c.JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, "invalid params"), fiber.StatusBadRequest))
	}
	id, err := strconv.Atoi(strID)
	if err != nil {
		return c.JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	req.ReceiverID = uint(id)
	userID := c.Locals("userID").(uint)
	req.SenderID = userID
	if err := h.service.AddFriends(c.Context(), req); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, "friend request sent successfully"))
}
