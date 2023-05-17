package request_body

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
)

type CreateSharedShoppingList struct {
	ShoppingListId *uuid.UUID `json:"shoppingListId,omitempty"`
	Name           *string    `json:"name,omitempty"`
}

type GetShoppingList struct {
	Key *string `json:"key,omitempty"`
}

type SetShoppingListName struct {
	Name *string `json:"name,omitempty"`
}

type SetShoppingList struct {
	Purchases   []common_body.Purchase `json:"purchases"`
	RecipeNames map[string]string      `json:"recipeNames"`
	LastVersion *int32                 `json:"lastVersion,omitempty"`
}

func Purchases(dtos []common_body.Purchase) []*api.Purchase {
	purchases := make([]*api.Purchase, len(dtos))
	for i, dto := range dtos {
		multiplier := 0
		if dto.Multiplier != nil && *dto.Multiplier > 0 {
			multiplier = *dto.Multiplier
		}

		amount := 0
		if dto.Amount != nil && *dto.Amount > 0 {
			amount = *dto.Amount
		}

		measureUnit := ""
		if dto.MeasureUnit != nil {
			measureUnit = *dto.MeasureUnit
		}

		recipeId := ""
		if dto.RecipeId != nil {
			recipeId = dto.RecipeId.String()
		}

		purchases[i] = &api.Purchase{
			Id:          dto.Id.String(),
			Name:        dto.Name,
			Multiplier:  int32(multiplier),
			Purchased:   dto.Purchased,
			Amount:      int32(amount),
			MeasureUnit: measureUnit,
			RecipeId:    recipeId,
		}
	}
	return purchases
}
