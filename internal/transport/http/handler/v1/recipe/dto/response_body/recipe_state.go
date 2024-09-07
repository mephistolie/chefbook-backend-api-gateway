package response_body

import (
	common "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

type RecipeState struct {
	Id string `json:"id"`

	Version int32 `json:"version"`

	Translations map[string][]string `json:"translations,omitempty"`

	Rating Rating `json:"rating"`

	Tags        []string `json:"tags,omitempty"`
	Collections []string `json:"collections,omitempty"`
	IsFavourite bool     `json:"favourite,omitempty"`
}

type GetRecipeBookResponse struct {
	Recipes                 []RecipeState                    `json:"recipes"`
	Tags                    map[string]Tag                   `json:"tags"`
	TagGroups               map[string]string                `json:"tagGroups"`
	Collections             []Collection                     `json:"collections"`
	IsEncryptedVaultEnabled bool                             `json:"isEncryptedVaultEnabled"`
	ProfilesInfo            map[string]common.ProfileMinInfo `json:"profilesInfo"`
}

func GetRecipeBook(response *api.GetRecipeBookResponse) GetRecipeBookResponse {
	return GetRecipeBookResponse{
		Recipes:                 newRecipeStates(response.Recipes),
		Tags:                    newTags(response.Tags),
		TagGroups:               common.NonNilStringMap(response.TagGroups),
		Collections:             newCollections(response.Collections),
		IsEncryptedVaultEnabled: response.HasEncryptedVault,
		ProfilesInfo:            newProfilesInfo(response.ProfilesInfo),
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

		Version: response.Version,

		Translations: newRecipeTranslations(response.Translations),

		Rating: Rating{
			Index: response.Rating,
			Score: response.Score,
			Votes: response.Votes,
		},

		Tags:        response.Tags,
		Collections: response.Collections,
		IsFavourite: response.IsFavourite,
	}
}
