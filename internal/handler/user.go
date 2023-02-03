package handler

import (
	"errors"
	"fiberun/internal/model"
	"fiberun/module/auth"
	"fiberun/module/base"
	"github.com/gookit/validate"
	"gorm.io/gorm"
)

type UserHandler struct{}

func (UserHandler) CreateUser(user *model.User) error {
	return base.GetDB().Create(user).Error
}

func (UserHandler) FindOne(find model.User) (model.User, bool) {
	user := model.User{}
	err := base.GetDB().Where(&find).First(&user).Error
	ok := !errors.Is(err, gorm.ErrRecordNotFound)
	return user, ok
}

func (UserHandler) ValidateParams(user *model.User) error {
	matcher := validate.Struct(user)
	if ok := matcher.Validate(); !ok {
		return matcher.Errors
	}
	return nil
}

func (UserHandler) EncryptHash(user *model.User) error {
	password, err := auth.EncryptHash(user.Password)
	if err == nil {
		user.Password = password
	}
	return err
}

func (UserHandler) CompareHash(user *model.User, pwd string) error {
	return auth.CompareHash(user.Password, pwd)
}

func (UserHandler) GenerateToken(user *model.User) (token string, err error) {
	token, err = auth.GenerateToken(user.Username, user.Password)
	return
}
