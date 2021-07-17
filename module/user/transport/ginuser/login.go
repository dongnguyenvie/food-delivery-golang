package ginuser

import (
	"net/http"
	"nolan/g05-food-delivery/common"
	"nolan/g05-food-delivery/component/appctx"
	"nolan/g05-food-delivery/component/hasher"
	"nolan/g05-food-delivery/component/tokenprovider/jwt"
	userbiz "nolan/g05-food-delivery/module/user/biz"
	usermodel "nolan/g05-food-delivery/module/user/model"
	userstorage "nolan/g05-food-delivery/module/user/storage"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetMaiDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBiz(store, md5, tokenProvider, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
