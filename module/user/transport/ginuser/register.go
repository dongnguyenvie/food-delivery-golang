package ginuser

import (
	"net/http"
	"nolan/g05-food-delivery/common"
	"nolan/g05-food-delivery/component/appctx"
	usermodel "nolan/g05-food-delivery/module/user/model"

	userbiz "nolan/g05-food-delivery/module/user/biz"
	userstorage "nolan/g05-food-delivery/module/user/storage"

	"nolan/g05-food-delivery/component/hasher"

	"github.com/gin-gonic/gin"
)

func Register(appctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appctx.GetMaiDBConnection()

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.RegisterUser(c, &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
