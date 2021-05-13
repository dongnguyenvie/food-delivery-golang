package main

import (
	"log"
	"net/http"
	"nolan/g05-food-delivery/component/appctx"
	"nolan/g05-food-delivery/middleware"
	"nolan/g05-food-delivery/module/restaurant/transport/ginrestaurant"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//`name` varchar(50) NOT NULL,
//`addr` varchar(255) NOT NULL,

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"` // tag
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	appContext := appctx.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// POST /restaurants
	v1 := r.Group("/v1")

	restaurants := v1.Group("/restaurants")

	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		var data Restaurant

		db.Where("id = ?", id).First(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

	restaurants.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		var data RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		db.Where("id = ?", id).Updates(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	r.Run()

	//newRestaurant := Restaurant{Name: "Tani", Addr: "9 Pham Van Hai"}
	//
	//if err := db.Create(&newRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println("New id:", newRestaurant.Id)

	//var myRestaurant Restaurant
	//
	//if err := db.Where("id = ?", 3).First(&myRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(myRestaurant)
	//
	//newName := "200Lab"
	//updateData := RestaurantUpdate{Name: &newName}
	//
	//if err := db.Where("id = ?", 3).Updates(&updateData).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(myRestaurant)
	//
	//if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 1).Delete(nil).Error; err != nil {
	//	log.Println(err)
	//}
}
