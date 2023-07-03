package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"time"
)

type CreateRecipe struct {
	RecipeId string `json:"recipeId"`
	Version  int32  `json:"version"`
}

type Recipe struct {
	Id   string `json:"recipeId"`
	Name string `json:"name"`

	OwnerId     string  `json:"ownerId"`
	OwnerName   *string `json:"ownerName,omitempty"`
	OwnerAvatar *string `json:"ownerAvatar,omitempty"`

	IsOwned     bool   `json:"owned"`
	IsSaved     bool   `json:"saved"`
	Visibility  string `json:"visibility"`
	IsEncrypted bool   `json:"encrypted"`

	Language    string  `json:"language"`
	Description *string `json:"description,omitempty"`
	Preview     *string `json:"preview,omitempty"`

	CreationTimestamp time.Time `json:"creationTimestamp"`
	UpdateTimestamp   time.Time `json:"updateTimestamp"`
	Version           int32     `json:"version"`

	Rating float32 `json:"rating"`
	Score  *int32  `json:"score,omitempty"`
	Votes  int32   `json:"votes"`

	Tags        []string `json:"tags"`
	Categories  []string `json:"categories,omitempty"`
	IsFavourite bool     `json:"favourite"`

	Servings *int32 `json:"servings,omitempty"`
	Time     *int32 `json:"time,omitempty"`

	Calories       *int32                      `json:"calories,omitempty"`
	Macronutrients *common_body.Macronutrients `json:"macronutrients,omitempty"`

	Ingredients []common_body.IngredientItem `json:"ingredients"`
	Cooking     []common_body.CookingItem    `json:"cooking"`
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

		OwnerId:     response.OwnerId,
		OwnerName:   response.OwnerName,
		OwnerAvatar: response.OwnerAvatar,

		IsOwned:     response.IsOwned,
		IsSaved:     response.IsSaved,
		Visibility:  response.Visibility,
		IsEncrypted: response.IsEncrypted,

		Language:    response.Language,
		Description: response.Description,
		Preview:     response.Preview,

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

		Calories:       response.Calories,
		Macronutrients: newMacronutrients(response.Macronutrients),

		Ingredients: newIngredients(response.Ingredients),
		Cooking:     newCooking(response.Cooking),
	}
}

type UpdateRecipe struct {
	Version int32 `json:"version"`
}
