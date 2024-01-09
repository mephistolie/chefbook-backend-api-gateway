package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

func newCooking(response []*api.CookingItem) []common_body.CookingItem {
	ingredients := make([]common_body.CookingItem, len(response))
	for id, item := range response {
		ingredients[id] = newCookingItem(item)
	}
	return ingredients
}

func newCookingItem(response *api.CookingItem) common_body.CookingItem {
	return common_body.CookingItem{
		Id:       response.Id,
		Text:     response.Text,
		Type:     response.Type,
		Time:     response.Time,
		RecipeId: response.RecipeId,
	}
}
