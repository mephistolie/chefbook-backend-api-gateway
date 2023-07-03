package response_body

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

type RecipeState struct {
	Id string `json:"recipeId"`

	OwnerName   *string `json:"ownerName,omitempty"`
	OwnerAvatar *string `json:"ownerAvatar,omitempty"`

	Preview *string `json:"preview,omitempty"`

	Version int32 `json:"version"`

	Rating float32 `json:"rating"`
	Score  *int32  `json:"score,omitempty"`
	Votes  int32   `json:"votes"`

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
	recipes := make([]RecipeState, len(response))
	for id, recipe := range response {
		recipes[id] = newRecipeState(recipe)
	}
	return recipes
}

func newRecipeState(response *api.RecipeState) RecipeState {
	return RecipeState{
		Id: response.RecipeId,

		OwnerName:   response.OwnerName,
		OwnerAvatar: response.OwnerAvatar,

		Version: response.Version,

		Rating: response.Rating,
		Score:  response.Score,
		Votes:  response.Votes,

		Categories:  response.Categories,
		IsFavourite: response.IsFavourite,
	}
}
