package userbiz

import (
	"context"
	"nolan/g05-food-delivery/common"
	usermodel "nolan/g05-food-delivery/module/user/model"
)

type ListStorage interface {
	ListDataWithCondition(context context.Context, filter *usermodel.Filter, paging *common.Paging) ([]usermodel.User, error)
}

type ListBiz struct {
	store ListStorage
}

func NewListBiz(store ListStorage) *ListBiz {
	return &ListBiz{store: store}
}

func (biz *ListBiz) ListUser(context context.Context, filter *usermodel.Filter, paging *common.Paging) ([]usermodel.User, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}
