package usersevice

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
	"os"
	"time"
)

func (s *userService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	var (
		user model.User
		resp dto.LoginResponse
		err  error
	)
	user, err = s.userStorage.OneRecord(ctx, "username = ?", req.Username)
	if err != nil {
		return resp, errors.New(constant.FailedUserNameOrPass)
	}
	if pwd := s.utils.PasswordControl(user.Password, req.Password); !pwd {
		return resp, errors.New(constant.FailedUserNameOrPass)
	}
	token, err := s.createJwt(user.ID)
	if err != nil {
		return resp, err
	}
	if err = s.userStorage.ChangeActive(ctx, user.ID, true); err != nil {
		return resp, err
	}
	resp.ID = user.ID
	resp.Username = user.Username
	resp.Phone = user.Phone
	resp.Token = token

	return resp, nil
}

func (s *userService) createJwt(userID uint) (string, error) {
	claims := &util.JwtCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return tokenSigned, err
	}
	return tokenSigned, nil
}
