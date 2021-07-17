package restaurantstorage

import (
	"context"
	common "nolan/g05-food-delivery/common"
	restaurantmodel "nolan/g05-food-delivery/module/restaurant/model"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
