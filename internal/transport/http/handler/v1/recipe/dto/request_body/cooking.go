package request_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

func newCooking(request []common_body.CookingItem) []*api.CookingItem {
	cooking := make([]*api.CookingItem, len(request))
	for id, ingredient := range request {
		cooking[id] = newCookingItem(ingredient)
	}
	return cooking
}

func newCookingItem(response common_body.CookingItem) *api.CookingItem {
	return &api.CookingItem{
		Id:       response.Id,
		Text:     response.Text,
		Type:     response.Type,
		Time:     response.Time,
		RecipeId: response.RecipeId,
	}
}
