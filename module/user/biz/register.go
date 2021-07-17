package userbiz

import (
	"context"
	"errors"
	"fmt"
	"nolan/g05-food-delivery/common"
	usermodel "nolan/g05-food-delivery/module/user/model"
)

type RegisterStorage interface {
	FindDataWithCondition(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	Create(context context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type RegisterBiz struct {
	store  RegisterStorage
	hasher Hasher
}

func NewRegisterBiz(store RegisterStorage, hasher Hasher) *RegisterBiz {
	return &RegisterBiz{store: store, hasher: hasher}
}

func (biz *RegisterBiz) RegisterUser(ctx context.Context, data *usermodel.UserCreate) error {

	userFound, _ := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"email": data.Email})
	if userFound != nil {
		return usermodel.ErrEmailExisted
	}

	if data.Email == "" {
		return errors.New("email is not empty")
	}
	salt := common.GenSalt(50)

	fmt.Print("salt: ", salt)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
