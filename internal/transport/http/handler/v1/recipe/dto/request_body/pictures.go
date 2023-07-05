package request_body

import "github.com/google/uuid"

type GenerateRecipePicturesUploadLinks struct {
	PicturesCount int32 `json:"picturesCount" binding:"required"`
}

type SetRecipePictures struct {
	Preview *uuid.UUID                 `json:"preview"`
	Cooking *map[uuid.UUID][]uuid.UUID `json:"cooking"`
	Version *int32                     `json:"version"`
}
