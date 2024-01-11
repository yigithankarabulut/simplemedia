package usersevice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
)

func (s *userService) UpdateProfilePicture(ctx context.Context, req dto.UpdatePictureRequest) error {
	var (
		user *model.User
		err  error
	)
	if user, err = s.userStorage.GetByID(ctx, req.UserID); err != nil {
		return errors.New("user not found")
	}
	if req.Username != user.Username {
		return errors.New("username mismatch")
	}
	user.ProfilePicture = req.PictureUrl
	if err := s.userStorage.Update(ctx, user); err != nil {
		return err
	}
	return nil
}
