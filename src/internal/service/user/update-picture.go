package usersevice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *userService) UpdateProfilePicture(ctx context.Context, req dto.UpdatePictureRequest) error {
	var (
		user *model.User
		err  error
	)
	if user, err = s.userStorage.GetByID(ctx, req.UserID); err != nil {
		return errors.New(constant.FailedUserNotFound)
	}
	if req.Username != user.Username {
		return errors.New(constant.FailedUNameMismatch)
	}
	user.ProfilePicture = req.PictureUrl
	if err := s.userStorage.Update(ctx, user); err != nil {
		return errors.New(fmt.Sprintf(constant.FailedSavePicture, err.Error()))
	}
	return nil
}
