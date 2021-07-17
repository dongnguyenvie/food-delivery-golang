package ginuser

import (
	"net/http"
	"nolan/g05-food-delivery/common"
	"nolan/g05-food-delivery/component/appctx"
	userbiz "nolan/g05-food-delivery/module/user/biz"
	usermodel "nolan/g05-food-delivery/module/user/model"
	userstorage "nolan/g05-food-delivery/module/user/storage"

	"github.com/gin-gonic/gin"
)

func List(appctx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		var pagingData *common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		db := appctx.GetMaiDBConnection()

		store := userstorage.NewSQLStore(db)

		biz := userbiz.NewListBiz(store)

		var filter *usermodel.Filter

		result, err := biz.ListUser(c, filter, pagingData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))

	}
}
