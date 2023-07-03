package response_body

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"time"
)

type RecipeInfo struct {
	Id   string `json:"recipeId"`
	Name string `json:"name"`

	OwnerId     string  `json:"ownerId"`
	OwnerName   *string `json:"ownerName,omitempty"`
	OwnerAvatar *string `json:"ownerAvatar,omitempty"`

	IsOwned     bool   `json:"owned"`
	IsSaved     bool   `json:"saved"`
	Visibility  string `json:"visibility"`
	IsEncrypted bool   `json:"encrypted,omitempty"`

	Language string  `json:"language"`
	Preview  *string `json:"preview,omitempty"`

	CreationTimestamp time.Time `json:"creationTimestamp"`
	UpdateTimestamp   time.Time `json:"updateTimestamp"`
	Version           int32     `json:"version"`

	Rating float32 `json:"rating"`
	Score  *int32  `json:"score,omitempty"`
	Votes  int32   `json:"votes"`

	Tags        []string `json:"tags,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	IsFavourite bool     `json:"favourite"`

	Servings *int32 `json:"servings,omitempty"`
	Time     *int32 `json:"time,omitempty"`

	Calories *int32 `json:"calories,omitempty"`
}

type GetRecipesResponse struct {
	Recipes    []RecipeInfo        `json:"recipes"`
	Tags       map[string]Tag      `json:"tags"`
	Categories map[string]Category `json:"categories"`
}

func GetRecipes(response *api.GetRecipesResponse) GetRecipesResponse {
	return GetRecipesResponse{
		Recipes:    newRecipeInfos(response.Recipes),
		Tags:       newTags(response.Tags),
		Categories: newCategoriesMap(response.Categories),
	}
}

func newRecipeInfos(response []*api.RecipeInfo) []RecipeInfo {
	recipes := make([]RecipeInfo, len(response))
	for id, recipe := range response {
		recipes[id] = newRecipeInfo(recipe)
	}
	return recipes
}

func newRecipeInfo(response *api.RecipeInfo) RecipeInfo {
	return RecipeInfo{
		Id:   response.RecipeId,
		Name: response.Name,

		OwnerId:     response.OwnerId,
		OwnerName:   response.OwnerName,
		OwnerAvatar: response.OwnerAvatar,

		IsOwned:     response.IsOwned,
		IsSaved:     response.IsSaved,
		Visibility:  response.Visibility,
		IsEncrypted: response.IsEncrypted,

		Language: response.Language,
		Preview:  response.Preview,

		CreationTimestamp: response.CreationTimestamp.AsTime(),
		UpdateTimestamp:   response.UpdateTimestamp.AsTime(),
		Version:           response.Version,

		Rating: response.Rating,
		Score:  response.Score,
		Votes:  response.Votes,

		Tags:        response.Tags,
		Categories:  response.Categories,
		IsFavourite: response.IsFavourite,

		Servings: response.Servings,
		Time:     response.Time,

		Calories: response.Calories,
	}
}
