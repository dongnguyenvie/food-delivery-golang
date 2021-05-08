package ginrestaurant

import (
	"net/http"
	"nolan/g05-food-delivery/common"
	appctx "nolan/g05-food-delivery/component/appctx"
	restaurantbiz "nolan/g05-food-delivery/module/restaurant/biz"
	restaurantmodel "nolan/g05-food-delivery/module/restaurant/model"
	restaurantstorage "nolan/g05-food-delivery/module/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		pagingData.Fullfill()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var result []restaurantmodel.Restaurant

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessReponse(result, pagingData, filter))
	}
}
