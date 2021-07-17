package userstorage

import (
	"context"
	"nolan/g05-food-delivery/common"
	usermodel "nolan/g05-food-delivery/module/user/model"
)

func (s *sqlStore) Delete(context context.Context, id string) error {
	if err := s.db.Table(usermodel.User{}.TableName()).
		Find("id = ?", id).
		Updates(map[string]int{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
