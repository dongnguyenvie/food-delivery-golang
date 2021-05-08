package main

import (
	"log"
	"net/http"
	"nolan/g05-food-delivery/component/appctx"
	"nolan/g05-food-delivery/module/restaurant/transport/ginrestaurant"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return "restaurants" }

func main() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=true&loc=Local"
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()

	appContext := appctx.NewAppContext(db)

	if err != nil {
		log.Fatalln(err)
	}

	// newRestaurant := Restaurant{Name: "Tani", Addr: "9 pham vang dong"}
	// if err := db.Create(&newRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }
	// log.Println(newRestaurant.Id)

	// var myRestaurant Restaurant
	// if err := db.Where("id = ?", 3).First(&myRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }

	// myRestaurant.Name = "name_changed"
	// if err := db.Where("id = ?", 3).Updates(&myRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }
	// log.Println(myRestaurant)

	// if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 2).Delete(nil).Error; err != nil {
	// 	log.Println(err)
	// }

	// newName := ""
	// updateData := RestaurantUpdate{Name: &newName}

	// myRestaurant.Name = "name_changed"
	// if err := db.Where("id = ?", 3).Updates(&updateData).Error; err != nil {
	// 	log.Println(err)
	// }
	log.Println(db)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")

	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

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
			"message": &data,
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
