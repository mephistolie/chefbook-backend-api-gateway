package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

type RecipePictureUpload struct {
	PictureLink string            `json:"pictureLink"`
	UploadLink  string            `json:"uploadLink"`
	FormData    map[string]string `json:"formData"`
	MaxSize     int64             `json:"maxSize"`
}

type SetRecipePictures struct {
	Pictures common_body.RecipePictures `json:"pictures"`
	Version  int32                      `json:"version"`
}

func newPicturesResponse(pictures *api.RecipePictures) *common_body.RecipePictures {
	if pictures == nil {
		return nil
	}
	res := common_body.NewPicturesResponse(pictures)

	return &res
}
