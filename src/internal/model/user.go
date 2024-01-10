package model

import (
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user/dto"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `json:"username" gorm:"size:50"`
	Email          string `json:"email" gorm:"size:50"`
	Password       string `json:"password"`
	Phone          string `json:"phone" gorm:"uniqueIndex;type:varchar(20)"`
	ProfilePicture string `json:"profile_picture"`
	IsActive       bool   `json:"is_active"`
}

func (u *User) Mapper(req dto.RegisterRequest) error {
	var utils util.Util

	pass, err := utils.GeneratePassword(req.Password)
	if err != nil {
		return err
	}
	u.Username = req.Username
	u.Email = req.Email
	u.Password = pass
	u.Phone = req.Phone
	if req.Picture != "" {
		u.ProfilePicture = req.Picture
	}
	return nil
}
