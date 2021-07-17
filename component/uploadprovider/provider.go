package uploadprovider

import (
	"context"
	"nolan/g05-food-delivery/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
