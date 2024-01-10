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
		return c.JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, "invalid params"), fiber.StatusBadRequest))
	}
	id, err := strconv.Atoi(strID)
	if err != nil {
		return c.JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	userID := c.Locals("userID").(uint)
	if userID == 0 {
		return c.JSON(h.util.BasicError(constant.UnknowError, fiber.StatusUnauthorized))
	}
	friend.SenderID = uint(id)
	friend.ReceiverID = userID
	if err := h.service.AcceptFriends(c.Context(), friend); err != nil {
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.JSON(h.util.Response(fiber.StatusOK, "friend request accepted successfully"))
}
