package userstorage

import (
	"context"
	"nolan/g05-food-delivery/common"
	usermodel "nolan/g05-food-delivery/module/user/model"
)

func (s *sqlStore) ListDataWithCondition(context context.Context, filter *usermodel.Filter, paging *common.Paging) ([]usermodel.User, error) {
	var result []usermodel.User
	db := s.db.Table(usermodel.User{}.TableName()).Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		// uid := common.NewUID(uint32(last.Id))
		paging.NextCursor = last.FakeId.String()
	}
	return result, nil
}
