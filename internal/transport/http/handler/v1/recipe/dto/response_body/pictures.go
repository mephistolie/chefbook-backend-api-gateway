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
	Version int32 `json:"version"`
}

func newPictures(pictures *api.RecipePictures) *common_body.RecipePictures {
	if pictures == nil {
		return nil
	}

	cooking := make(map[string][]string)
	for stepId, stepPictures := range pictures.Cooking {
		cooking[stepId] = stepPictures.Pictures
	}

	return &common_body.RecipePictures{
		Preview: pictures.Preview,
		Cooking: cooking,
	}
}
