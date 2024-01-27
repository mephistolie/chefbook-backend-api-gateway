package common_body

import api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"

type RecipePictures struct {
	Preview *string             `json:"preview,omitempty"`
	Cooking map[string][]string `json:"cooking,omitempty"`
}

func NewPicturesRequest(pictures RecipePictures) *api.RecipePictures {
	cooking := make(map[string]*api.StepPictures)
	for stepId, stepPictures := range pictures.Cooking {
		cooking[stepId] = &api.StepPictures{Pictures: stepPictures}
	}

	return &api.RecipePictures{
		Preview: pictures.Preview,
		Cooking: cooking,
	}
}

func NewPicturesResponse(pictures *api.RecipePictures) RecipePictures {
	cooking := make(map[string][]string)
	for stepId, stepPictures := range pictures.Cooking {
		cooking[stepId] = stepPictures.Pictures
	}

	return RecipePictures{
		Preview: pictures.Preview,
		Cooking: cooking,
	}
}
