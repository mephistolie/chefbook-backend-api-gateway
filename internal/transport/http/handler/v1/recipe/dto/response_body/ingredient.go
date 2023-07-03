package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

func newIngredients(response []*api.IngredientItem) []common_body.IngredientItem {
	ingredients := make([]common_body.IngredientItem, len(response))
	for id, ingredient := range response {
		ingredients[id] = newIngredient(ingredient)
	}
	return ingredients
}

func newIngredient(response *api.IngredientItem) common_body.IngredientItem {
	return common_body.IngredientItem{
		Id:       response.Id,
		Text:     response.Text,
		Type:     response.Type,
		Unit:     response.Unit,
		Amount:   response.Amount,
		RecipeId: response.RecipeId,
	}
}
