package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list"
)

func (r *Router) initShoppingListsRoutes(api *gin.RouterGroup) {
	shoppingListGroup := api.Group("/shopping-lists", r.authMiddleware.AuthorizeUser)
	{
		shoppingListGroup.GET("", r.handler.ShoppingList.GetShoppingLists)

		shoppingListGroup.POST("", r.handler.ShoppingList.CreateSharedShoppingList)
		shoppingListGroup.GET("/personal", r.handler.ShoppingList.GetPersonalShoppingList)
		shoppingListGroup.GET(fmt.Sprintf("/:%s", shopping_list.ParamShoppingListId), r.handler.ShoppingList.GetShoppingList)
		shoppingListGroup.PUT(fmt.Sprintf("/:%s/name", shopping_list.ParamShoppingListId), r.handler.ShoppingList.SetShoppingListName)
		shoppingListGroup.PUT(fmt.Sprintf("/:%s", shopping_list.ParamShoppingListId), r.handler.ShoppingList.SetShoppingList)
		shoppingListGroup.PATCH(fmt.Sprintf("/:%s", shopping_list.ParamShoppingListId), r.handler.ShoppingList.AddToShoppingList)
		shoppingListGroup.DELETE(fmt.Sprintf("/:%s", shopping_list.ParamShoppingListId), r.handler.ShoppingList.DeleteSharedShoppingList)

		shoppingListGroup.GET(fmt.Sprintf("/:%s/link", shopping_list.ParamShoppingListId), r.handler.ShoppingList.GetSharedShoppingListLink)
		shoppingListGroup.GET(fmt.Sprintf("/:%s/users", shopping_list.ParamShoppingListId), r.handler.ShoppingList.GetShoppingListUsers)
		shoppingListGroup.POST(fmt.Sprintf("/:%s/users", shopping_list.ParamShoppingListId), r.handler.ShoppingList.JoinShoppingList)
		shoppingListGroup.DELETE(fmt.Sprintf("/:%s/users", shopping_list.ParamShoppingListId), r.handler.ShoppingList.DeleteUserFromShoppingList)
	}
}
