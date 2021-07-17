package restaurantstorage

import (
	"context"
	"nolan/g05-food-delivery/common"
	restaurantmodel "nolan/g05-food-delivery/module/restaurant/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
