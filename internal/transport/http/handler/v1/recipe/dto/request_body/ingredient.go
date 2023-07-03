package request_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

func newIngredients(request []common_body.IngredientItem) []*api.IngredientItem {
	ingredients := make([]*api.IngredientItem, len(request))
	for id, ingredient := range request {
		ingredients[id] = newIngredient(ingredient)
	}
	return ingredients
}

func newIngredient(response common_body.IngredientItem) *api.IngredientItem {
	return &api.IngredientItem{
		Id:       response.Id,
		Text:     response.Text,
		Type:     response.Type,
		Unit:     response.Unit,
		Amount:   response.Amount,
		RecipeId: response.RecipeId,
	}
}
