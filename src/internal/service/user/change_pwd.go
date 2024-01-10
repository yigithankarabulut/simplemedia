package usersevice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *userService) ChangePassword(ctx context.Context, userID uint, req dto.ChangePwdRequest) error {
	user, err := s.userStorage.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	if ok := s.utils.PasswordControl(user.Password, req.OldPassword); !ok {
		return errors.New(fmt.Sprintf(constant.FailedChangePwd, "password is wrong"))
	}
	if req.NewPassword == req.OldPassword {
		return errors.New(fmt.Sprintf(constant.FailedChangePwd, "new password can not be the same as the old password"))
	}
	user.Password, err = s.utils.GeneratePassword(req.NewPassword)
	if err != nil {
		return err
	}
	if err := s.userStorage.Update(ctx, user); err != nil {
		return err
	}
	return nil
}
