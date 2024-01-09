package response_body

import (
	"github.com/mephistolie/chefbook-backend-common/log"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

type RecipeState struct {
	Id string `json:"id"`

	Owner *ProfileInfo `json:"owner,omitempty"`

	Version int32 `json:"version"`

	Translations []string `json:"translations"`

	Rating Rating `json:"rating"`

	Categories  []string `json:"categories,omitempty"`
	IsFavourite bool     `json:"favourite"`
}

type GetRecipeBookResponse struct {
	Recipes    []RecipeState  `json:"recipes"`
	Tags       map[string]Tag `json:"tags"`
	Categories []Category     `json:"categories"`
}

func GetRecipeBook(response *api.GetRecipeBookResponse) GetRecipeBookResponse {
	return GetRecipeBookResponse{
		Recipes:    newRecipeStates(response.Recipes),
		Tags:       newTags(response.Tags),
		Categories: newCategories(response.Categories),
	}
}

func newRecipeStates(response []*api.RecipeState) []RecipeState {
	log.Debugf("got %d recipes by request", len(response))
	recipes := make([]RecipeState, len(response))
	for id, recipe := range response {
		recipes[id] = newRecipeState(recipe)
	}
	return recipes
}

func newRecipeState(response *api.RecipeState) RecipeState {
	var owner *ProfileInfo
	if response.OwnerName != nil || response.OwnerAvatar != nil {
		owner = &ProfileInfo{
			Name:   response.OwnerName,
			Avatar: response.OwnerAvatar,
		}
	}

	return RecipeState{
		Id: response.RecipeId,

		Owner: owner,

		Version: response.Version,

		Translations: response.Translations,

		Rating: Rating{
			Index: response.Rating,
			Score: response.Score,
			Votes: response.Votes,
		},

		Categories:  response.Categories,
		IsFavourite: response.IsFavourite,
	}
}
