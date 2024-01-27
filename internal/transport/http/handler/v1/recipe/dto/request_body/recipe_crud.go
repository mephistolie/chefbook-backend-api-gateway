package request_body

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

type RecipeInput struct {
	Id   *string `json:"recipeId"`
	Name string  `json:"name" binding:"required"`

	Visibility  string `json:"visibility,omitempty"`
	IsEncrypted bool   `json:"encrypted,omitempty"`

	Language    *string `json:"language,omitempty"`
	Description *string `json:"description,omitempty"`

	Version *int32 `json:"version,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	Servings *int32 `json:"servings,omitempty"`
	Time     *int32 `json:"time,omitempty"`

	Calories       *int32                      `json:"calories,omitempty"`
	Macronutrients *common_body.Macronutrients `json:"macronutrients,omitempty"`

	Ingredients []common_body.IngredientItem `json:"ingredients" binding:"required"`
	Cooking     []common_body.CookingItem    `json:"cooking" binding:"required"`
}

func RecipeInputRequest(body RecipeInput, userId uuid.UUID) *api.RecipeInput {
	var tags []string
	if body.Tags != nil {
		tags = *body.Tags
	}
	return &api.RecipeInput{
		RecipeId:       body.Id,
		UserId:         userId.String(),
		Name:           body.Name,
		Visibility:     body.Visibility,
		IsEncrypted:    body.IsEncrypted,
		Language:       body.Language,
		Description:    body.Description,
		Tags:           tags,
		Servings:       body.Servings,
		Time:           body.Time,
		Calories:       body.Calories,
		Macronutrients: newMacronutrients(body.Macronutrients),
		Ingredients:    newIngredients(body.Ingredients),
		Cooking:        newCooking(body.Cooking),
		Version:        body.Version,
	}
}
