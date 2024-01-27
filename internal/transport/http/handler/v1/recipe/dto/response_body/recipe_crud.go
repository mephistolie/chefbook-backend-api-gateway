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

	Owner common.ProfileInfo `json:"owner"`

	IsOwned     bool   `json:"owned"`
	IsSaved     bool   `json:"saved"`
	Visibility  string `json:"visibility"`
	IsEncrypted bool   `json:"encrypted"`

	Language     string                         `json:"language"`
	Translations map[string][]RecipeTranslation `json:"translations"`
	Description  *string                        `json:"description,omitempty"`

	CreationTimestamp time.Time `json:"creationTimestamp"`
	UpdateTimestamp   time.Time `json:"updateTimestamp"`
	Version           int32     `json:"version"`

	Rating Rating `json:"rating"`

	Tags        []string `json:"tags"`
	Categories  []string `json:"categories,omitempty"`
	IsFavourite bool     `json:"favourite"`

	Servings *int32 `json:"servings,omitempty"`
	Time     *int32 `json:"time,omitempty"`

	Calories       *int32                      `json:"calories,omitempty"`
	Macronutrients *common_body.Macronutrients `json:"macronutrients,omitempty"`

	Ingredients []common_body.IngredientItem `json:"ingredients"`
	Cooking     []common_body.CookingItem    `json:"cooking"`
	Pictures    *common_body.RecipePictures  `json:"pictures"`
}

type GetRecipeResponse struct {
	Recipe     Recipe              `json:"recipe"`
	Tags       map[string]Tag      `json:"tags"`
	Categories map[string]Category `json:"categories"`
}

func GetRecipe(response *api.GetRecipeResponse) GetRecipeResponse {
	return GetRecipeResponse{
		Recipe:     newRecipe(response.Recipe),
		Tags:       newTags(response.Tags),
		Categories: newCategoriesMap(response.Categories),
	}
}

func newRecipe(response *api.Recipe) Recipe {
	return Recipe{
		Id:   response.RecipeId,
		Name: response.Name,

		Owner: common.ProfileInfo{
			Id:     response.OwnerId,
			Name:   response.OwnerName,
			Avatar: response.OwnerAvatar,
		},

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
		Categories:  response.Categories,
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
