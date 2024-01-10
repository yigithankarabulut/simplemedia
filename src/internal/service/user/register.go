package usersevice

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
)

func (s *userService) Register(ctx context.Context, req dto.RegisterRequest) error {
	var (
		user model.User
		err  error
		tx   = s.userStorage.CreateTx()
	)
	if err = user.Mapper(req); err != nil {
		return err
	}
	defer tx.Rollback()

	if err = s.duplicateCheck(user.Username, user.Email, user.Phone); err != nil {
		return err
	}
	if err = s.userStorage.Insert(ctx, &user, tx); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (s *userService) duplicateCheck(username, email, phone string) error {
	var ctx context.Context
	if _, err := s.userStorage.OneRecordWithUserName(ctx, username); err == nil {
		return errors.New(constant.AlreadyExistsUsername)
	}
	if _, err := s.userStorage.OneRecordWithEmail(ctx, email); err == nil {
		return errors.New(constant.AlreadyExistsEmail)
	}
	if _, err := s.userStorage.OneRecordWithPhone(ctx, phone); err == nil {
		return errors.New(constant.AlreadyExistsPhone)
	}
	return nil
}
