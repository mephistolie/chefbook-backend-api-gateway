package response_body

import (
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
)

type ShoppingListInfo struct {
	Id      string `json:"shoppingListId"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type"`
	OwnerId string `json:"ownerId"`
}

func GetShoppingLists(response *api.GetShoppingListsResponse) []ShoppingListInfo {
	shoppingLists := make([]ShoppingListInfo, len(response.ShoppingLists))
	for i, shoppingList := range response.ShoppingLists {
		shoppingLists[i] = ShoppingListInfo{
			Id:      shoppingList.Id,
			Name:    shoppingList.Name,
			Type:    shoppingList.Type,
			OwnerId: shoppingList.OwnerId,
		}
	}
	return shoppingLists
}

type CreateShoppingList struct {
	Id string `json:"shoppingListId"`
}

type GetShoppingListBody struct {
	Id          string                 `json:"shoppingListId"`
	Name        string                 `json:"name,omitempty"`
	Type        string                 `json:"type"`
	OwnerId     string                 `json:"ownerId"`
	Purchases   []common_body.Purchase `json:"purchases"`
	RecipeNames map[string]string      `json:"recipeNames"`
	Version     int32                  `json:"version"`
}

func GetShoppingList(shoppingList *api.GetShoppingListResponse) GetShoppingListBody {
	dtos := make([]common_body.Purchase, len(shoppingList.Purchases))
	for i, purchase := range shoppingList.Purchases {
		id, err := uuid.Parse(purchase.Id)
		if err != nil {
			continue
		}

		var multiplierPtr *int = nil
		if purchase.Multiplier > 0 {
			multiplier := int(purchase.Multiplier)
			multiplierPtr = &multiplier
		}

		var amountPtr *int = nil
		if purchase.Amount > 0 {
			amount := int(purchase.Amount)
			amountPtr = &amount
		}

		var measureUnitPtr *string = nil
		if len(purchase.MeasureUnit) > 0 {
			measureUnit := purchase.MeasureUnit
			measureUnitPtr = &measureUnit
		}

		var recipeIdPtr *uuid.UUID = nil
		if recipeId, err := uuid.Parse(purchase.RecipeId); err == nil {
			recipeIdPtr = &recipeId
		}

		dtos[i] = common_body.Purchase{
			Id:          id,
			Name:        purchase.Name,
			Multiplier:  multiplierPtr,
			Purchased:   purchase.Purchased,
			Amount:      amountPtr,
			MeasureUnit: measureUnitPtr,
			RecipeId:    recipeIdPtr,
		}
	}
	return GetShoppingListBody{
		Id:          shoppingList.Id,
		Name:        shoppingList.Name,
		OwnerId:     shoppingList.OwnerId,
		Type:        shoppingList.Type,
		Purchases:   dtos,
		RecipeNames: shoppingList.RecipeNames,
		Version:     shoppingList.Version,
	}
}

type SetShoppingList struct {
	Version int32 `json:"version"`
}
