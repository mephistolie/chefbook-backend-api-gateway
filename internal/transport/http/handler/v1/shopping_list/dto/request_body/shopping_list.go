package request_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
)

type CreateSharedShoppingList struct {
	ShoppingListId *string `json:"shoppingListId,omitempty"`
	Name           *string `json:"name,omitempty"`
}

type SetShoppingListName struct {
	Name *string `json:"name,omitempty"`
}

type SetShoppingList struct {
	Purchases   []common_body.Purchase `json:"purchases"`
	LastVersion *int32                 `json:"lastVersion,omitempty"`
}

func Purchases(dtos []common_body.Purchase) []*api.Purchase {
	purchases := make([]*api.Purchase, len(dtos))
	for i, dto := range dtos {
		purchases[i] = &api.Purchase{
			Id:          dto.Id,
			Name:        dto.Name,
			Multiplier:  dto.Multiplier,
			Purchased:   dto.Purchased,
			Amount:      dto.Amount,
			MeasureUnit: dto.MeasureUnit,
			RecipeId:    dto.RecipeId,
		}
	}
	return purchases
}
