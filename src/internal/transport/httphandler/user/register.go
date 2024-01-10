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
		req  dto.RegisterRequest
		wg   sync.WaitGroup
		done = make(chan error, 1)
	)
	if err := h.util.Validate(c, &req); err != nil {
		return c.JSON(h.util.BasicError(fmt.Sprintf(constant.Validate, err.Error()), fiber.StatusBadRequest))
	}
	file, err := c.FormFile("picture")
	if err == nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := <-done
			if err != nil {
				return
			}
			filepath, err := h.util.SavePicture(file, req.Username, "profile-picture")
			if err != nil {
				done <- err
				return
			}
			if err := c.SaveFile(file, filepath); err != nil {
				done <- err
				return
			}
			req.PictureUrl = filepath
			done <- nil
		}()
	}
	if err := h.service.Register(c.Context(), &req, done); err != nil {
		close(done)
		return c.JSON(h.util.BasicError(err, fiber.StatusInternalServerError))
	}
	close(done)
	return c.JSON(h.util.Response(fiber.StatusOK, "User registered successfully"))
}
