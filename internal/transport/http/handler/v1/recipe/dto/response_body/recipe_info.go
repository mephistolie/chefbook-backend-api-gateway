package response_body

import (
	common "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"time"
)

type RecipeInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	OwnerId string `json:"ownerId"`

	IsOwned     bool   `json:"owned,omitempty"`
	IsSaved     bool   `json:"saved,omitempty"`
	Visibility  string `json:"visibility"`
	IsEncrypted bool   `json:"encrypted,omitempty"`

	Language     string   `json:"language"`
	Translations []string `json:"translations,omitempty"`
	Preview      *string  `json:"preview,omitempty"`

	CreationTimestamp time.Time `json:"creationTimestamp"`
	UpdateTimestamp   time.Time `json:"updateTimestamp"`
	Version           int32     `json:"version"`

	Rating Rating `json:"rating"`

	Tags        []string `json:"tags,omitempty"`
	Collections []string `json:"collections,omitempty"`
	IsFavourite bool     `json:"favourite,omitempty"`

	Servings *int32 `json:"servings,omitempty"`
	Time     *int32 `json:"time,omitempty"`

	Calories *int32 `json:"calories,omitempty"`
}

type GetRecipesResponse struct {
	Recipes      []RecipeInfo                     `json:"recipes"`
	Collections  map[string]CollectionInfo        `json:"collections"`
	Tags         map[string]Tag                   `json:"tags"`
	TagGroups    map[string]string                `json:"tagGroups"`
	ProfilesInfo map[string]common.ProfileMinInfo `json:"profilesInfo"`
}

func GetRecipes(response *api.GetRecipesResponse) GetRecipesResponse {
	return GetRecipesResponse{
		Recipes:      newRecipeInfos(response.Recipes),
		Collections:  newCollectionsMap(response.Collections),
		Tags:         newTags(response.Tags),
		TagGroups:    common.NonNilStringMap(response.TagGroups),
		ProfilesInfo: newProfilesInfo(response.ProfilesInfo),
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

		OwnerId: response.OwnerId,

		IsOwned:     response.IsOwned,
		IsSaved:     response.IsSaved,
		Visibility:  response.Visibility,
		IsEncrypted: response.IsEncrypted,

		Language:     response.Language,
		Translations: response.Translations,
		Preview:      response.Preview,

		CreationTimestamp: response.CreationTimestamp.AsTime(),
		UpdateTimestamp:   response.UpdateTimestamp.AsTime(),
		Version:           response.Version,

		Rating: Rating{
			Index: response.Rating,
			Score: response.Score,
			Votes: response.Votes,
		},

		Tags:        response.Tags,
		Collections: response.Collections,
		IsFavourite: response.IsFavourite,

		Servings: response.Servings,
		Time:     response.Time,

		Calories: response.Calories,
	}
}
