package main

import (
	"nolan/g05-food-delivery/component/appctx"
	"nolan/g05-food-delivery/middleware"

	"nolan/g05-food-delivery/module/restaurant/transport/ginrestaurant"
	"nolan/g05-food-delivery/module/upload/transport/ginupload"
	"nolan/g05-food-delivery/module/user/transport/ginuser"

	"nolan/g05-food-delivery/memcache"

	userstore "nolan/g05-food-delivery/module/user/storage"

	"github.com/gin-gonic/gin"
)

func setupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	userStore := userstore.NewSQLStore(appContext.GetMaiDBConnection())
	userCachingStore := memcache.NewUserCaching(memcache.NewCaching(), userStore)

	v1.POST("/login", ginuser.Login(appContext))
	v1.POST("/upload", ginupload.UploadImage(appContext))
	v1.GET("/profile", middleware.RequiredAuth(appContext, userCachingStore), ginuser.Profile(appContext))

	users := v1.Group("/users")
	{
		users.POST("", ginuser.Register(appContext))
		users.GET("", ginuser.List(appContext))
	}

	restaurants := v1.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
		restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

		// restaurants.PATCH("/:id", func(c *gin.Context) {
		// 	id, err := strconv.Atoi(c.Param("id"))

		// 	if err != nil {
		// 		c.JSON(http.StatusBadRequest, gin.H{
		// 			"error": err.Error(),
		// 		})

		// 		return
		// 	}

		// 	var data RestaurantUpdate

		// 	if err := c.ShouldBind(&data); err != nil {
		// 		c.JSON(http.StatusBadRequest, gin.H{
		// 			"error": err.Error(),
		// 		})

		// 		return
		// 	}

		// 	db.Where("id = ?", id).Updates(&data)

		// 	c.JSON(http.StatusOK, gin.H{
		// 		"data": data,
		// 	})
		// })

	}
}
