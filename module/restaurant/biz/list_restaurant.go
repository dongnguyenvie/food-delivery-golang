package restaurantbiz

import (
	"context"
	"nolan/g05-food-delivery/common"
	restaurantmodel "nolan/g05-food-delivery/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(context context.Context, filer *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(context, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}
