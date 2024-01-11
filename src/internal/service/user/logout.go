package usersevice

import (
	"context"
	"errors"
	"fmt"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *userService) Logout(ctx context.Context, userID uint) error {
	if err := s.userStorage.ChangeActive(ctx, userID, false); err != nil {
		return errors.New(fmt.Sprintf(constant.LogoutFailed, err.Error()))
	}
	return nil
}
