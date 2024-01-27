package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/common_body"
	common "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
)

type ShoppingListInfo struct {
	Id      string             `json:"shoppingListId"`
	Name    *string            `json:"name,omitempty"`
	Type    string             `json:"type"`
	Owner   common.ProfileInfo `json:"owner"`
	Version int32              `json:"version"`
}

func GetShoppingLists(response *api.GetShoppingListsResponse) []ShoppingListInfo {
	shoppingLists := make([]ShoppingListInfo, len(response.ShoppingLists))
	for i, shoppingList := range response.ShoppingLists {
		shoppingLists[i] = ShoppingListInfo{
			Id:   shoppingList.Id,
			Name: shoppingList.Name,
			Type: shoppingList.Type,
			Owner: common.ProfileInfo{
				Id:     shoppingList.Owner.Id,
				Name:   shoppingList.Owner.Name,
				Avatar: shoppingList.Owner.Avatar,
			},
			Version: shoppingList.Version,
		}
	}
	return shoppingLists
}

type CreateShoppingList struct {
	Id string `json:"shoppingListId"`
}

type GetShoppingListBody struct {
	Id          string                 `json:"shoppingListId"`
	Name        *string                `json:"name,omitempty"`
	Type        string                 `json:"type"`
	Owner       common.ProfileInfo     `json:"owner"`
	Purchases   []common_body.Purchase `json:"purchases"`
	RecipeNames map[string]string      `json:"recipeNames"`
	Version     int32                  `json:"version"`
}

func GetShoppingList(shoppingList *api.GetShoppingListResponse) GetShoppingListBody {
	dtos := make([]common_body.Purchase, len(shoppingList.Purchases))
	for i, purchase := range shoppingList.Purchases {
		dtos[i] = common_body.Purchase{
			Id:          purchase.Id,
			Name:        purchase.Name,
			Multiplier:  purchase.Multiplier,
			Purchased:   purchase.Purchased,
			Amount:      purchase.Amount,
			MeasureUnit: purchase.MeasureUnit,
			RecipeId:    purchase.RecipeId,
		}
	}
	return GetShoppingListBody{
		Id:   shoppingList.Id,
		Name: shoppingList.Name,
		Owner: common.ProfileInfo{
			Id:     shoppingList.Owner.Id,
			Name:   shoppingList.Owner.Name,
			Avatar: shoppingList.Owner.Avatar,
		},
		Type:        shoppingList.Type,
		Purchases:   dtos,
		RecipeNames: shoppingList.RecipeNames,
		Version:     shoppingList.Version,
	}
}

type SetShoppingList struct {
	Version int32 `json:"version"`
}
