package ginrestaurant

import (
	"net/http"
	"nolan/g05-food-delivery/common"
	"nolan/g05-food-delivery/component/appctx"
	restaurantbiz "nolan/g05-food-delivery/module/restaurant/biz"
	restaurantmodel "nolan/g05-food-delivery/module/restaurant/model"
	restaurantstorage "nolan/g05-food-delivery/module/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()

		//go func() {
		//	defer common.AppRecover()
		//
		//	arr := []int{}
		//	log.Println(arr[0])
		//}()

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
