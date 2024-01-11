package userhttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
	"sync"
)

func (h *Handler) Register(c *fiber.Ctx) error {
	var (
		req   dto.RegisterRequest
		wg    sync.WaitGroup
		errCh = make(chan error, 1)
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.Status(400).JSON(h.util.BasicError(
			fmt.Sprintf(constant.Validate, err.Error()),
			fiber.StatusBadRequest))
	}
	defer close(errCh)
	file, err := c.FormFile("picture")
	if err == nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := <-errCh
			if err != nil {
				return
			}
			filepath, err := h.util.SavePicture(file, "profile-picture", req.Username)
			if err != nil {
				errCh <- err
				return
			}
			if err := c.SaveFile(file, filepath); err != nil {
				errCh <- err
				return
			}
			req.PictureUrl = filepath
			errCh <- nil
		}()
	}
	if err := h.service.Register(c.Context(), &req, errCh); err != nil {
		return c.Status(500).JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	return c.Status(200).JSON(h.util.Response(fiber.StatusOK, "User created successfully"))
}
