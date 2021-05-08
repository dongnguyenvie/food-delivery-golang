package restaurantstorage

import (
	"context"
	"nolan/g05-food-delivery/common"
	restaurantmodel "nolan/g05-food-delivery/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where(("status in (1)"))

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
