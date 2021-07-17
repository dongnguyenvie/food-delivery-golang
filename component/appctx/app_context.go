package appctx

import (
	"nolan/g05-food-delivery/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider, secretKey string) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider, secretKey: secretKey}
}

func (ctx *appCtx) GetMaiDBConnection() *gorm.DB                  { return ctx.db }
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider { return ctx.uploadProvider }
func (ctx *appCtx) SecretKey() string                             { return ctx.secretKey }
