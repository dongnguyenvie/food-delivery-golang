package ginupload

import (
	"fmt"
	"nolan/g05-food-delivery/common"
	"nolan/g05-food-delivery/component/appctx"
	uploadbiz "nolan/g05-food-delivery/module/upload/biz"

	_ "image/jpeg"
	_ "image/png"

	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		// if err != nil {
		// 	panic(err)
		// }

		// c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename))

		// c.JSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
		// 	Id:        0,
		// 	Url:       fmt.Sprintf("http://localhost:8080/static/%s", fileHeader.Filename),
		// 	Width:     0,
		// 	Height:    0,
		// 	CloudName: "local",
		// 	Extension: "png",
		// }))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		fmt.Print("folder", folder)

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//imgStore := uploadstorage.NewSQLStore(db)
		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
