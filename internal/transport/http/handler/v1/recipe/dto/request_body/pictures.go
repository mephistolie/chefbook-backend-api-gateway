package request_body

import "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"

type GenerateRecipePicturesUploadLinks struct {
	PicturesCount int32 `json:"picturesCount" binding:"required"`
}

type SetRecipePictures struct {
	Pictures common_body.RecipePictures `json:"pictures"`
	Version  *int32                     `json:"version"`
}
