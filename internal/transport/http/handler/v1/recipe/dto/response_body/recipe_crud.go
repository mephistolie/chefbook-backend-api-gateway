package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	common "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"time"
)

type CreateRecipe struct {
	RecipeId string `json:"recipeId"`
	Version  int32  `json:"version"`
}

type Recipe struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	OwnerId string `json:"ownerId"`

	IsOwned     bool   `json:"owned,omitempty"`
	IsSaved     bool   `json:"saved,omitempty"`
	Visibility  string `json:"visibility"`
	IsEncrypted bool   `json:"encrypted,omitempty"`

	Language     string              `json:"language"`
	Translations map[string][]string `json:"translations,omitempty"`
	Description  *string             `json:"description,omitempty"`

	CreationTimestamp time.Time `json:"creationTimestamp"`
	UpdateTimestamp   time.Time `json:"updateTimestamp"`
	Version           int32     `json:"version"`

	Rating Rating `json:"rating"`

	Tags        []string `json:"tags,omitempty"`
	Collections []string `json:"collections,omitempty"`
	IsFavourite bool     `json:"favourite,omitempty"`

	Servings *int32 `json:"servings,omitempty"`
	Time     *int32 `json:"time,omitempty"`

	Calories       *int32                      `json:"calories,omitempty"`
	Macronutrients *common_body.Macronutrients `json:"macronutrients,omitempty"`

	Ingredients []common_body.IngredientItem `json:"ingredients"`
	Cooking     []common_body.CookingItem    `json:"cooking"`
	Pictures    *common_body.RecipePictures  `json:"pictures,omitempty"`
}

type GetRecipeResponse struct {
	Recipe       Recipe                           `json:"recipe"`
	Collections  map[string]CollectionInfo        `json:"collections"`
	Tags         map[string]Tag                   `json:"tags"`
	TagGroups    map[string]string                `json:"tagGroups"`
	ProfilesInfo map[string]common.ProfileMinInfo `json:"profilesInfo"`
}

func GetRecipe(response *api.GetRecipeResponse) GetRecipeResponse {
	return GetRecipeResponse{
		Recipe:       newRecipe(response.Recipe),
		Collections:  newCollectionsMap(response.Collections),
		Tags:         newTags(response.Tags),
		TagGroups:    common.NonNilStringMap(response.TagGroups),
		ProfilesInfo: newProfilesInfo(response.ProfilesInfo),
	}
}

func newRecipe(response *api.Recipe) Recipe {
	return Recipe{
		Id:   response.RecipeId,
		Name: response.Name,

		OwnerId: response.OwnerId,

		IsOwned:     response.IsOwned,
		IsSaved:     response.IsSaved,
		Visibility:  response.Visibility,
		IsEncrypted: response.IsEncrypted,

		Language:     response.Language,
		Translations: newRecipeTranslations(response.Translations),
		Description:  response.Description,

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

		Calories:       response.Calories,
		Macronutrients: newMacronutrients(response.Macronutrients),

		Ingredients: newIngredients(response.Ingredients),
		Cooking:     newCooking(response.Cooking),
		Pictures:    newPicturesResponse(response.Pictures),
	}
}

type UpdateRecipe struct {
	Version int32 `json:"version"`
}
