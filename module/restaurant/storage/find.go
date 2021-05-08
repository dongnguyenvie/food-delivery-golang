package restaurantstorage

import (
	"context"
	restaurantmodel "nolan/g05-food-delivery/module/restaurant/model"
)

func (s *sqlStore) FindDataWithCondition(context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
