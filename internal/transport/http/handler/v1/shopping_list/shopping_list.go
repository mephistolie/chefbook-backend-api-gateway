package shopping_list

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
)

// GetShoppingLists Swagger Documentation
//
//	@Summary		Get personal shopping list
//	@Description	Get personal shopping list
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	[]response_body.ShoppingListInfo
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/shopping-lists	[get]
func (h *Handler) GetShoppingLists(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetShoppingLists(c, &api.GetShoppingListsRequest{UserId: payload.UserId.String()})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetShoppingLists(res))
}

// CreateSharedShoppingList Swagger Documentation
//
//	@Summary		Create shared shopping list
//	@Description	Create shared shopping list
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.CreateSharedShoppingList	true	"Params"
//	@Success		200					{object}	response_body.GetShoppingListBody
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/shopping-lists	[post]
func (h *Handler) CreateSharedShoppingList(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.CreateSharedShoppingList
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}
	shoppingListId := ""
	if body.ShoppingListId != nil {
		shoppingListId = body.ShoppingListId.String()
	}
	name := ""
	if body.Name != nil {
		name = *body.Name
	}

	res, err := h.service.CreateSharedShoppingList(c, &api.CreateSharedShoppingListRequest{
		ShoppingListId:   shoppingListId,
		Name:             name,
		UserId:           payload.UserId.String(),
		SubscriptionPlan: payload.SubscriptionPlan,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.CreateShoppingList{Id: res.ShoppingListId})
}

// GetPersonalShoppingList Swagger Documentation
//
//	@Summary		Get personal shopping list
//	@Description	Get personal shopping list
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200							{object}	response_body.GetShoppingListBody
//	@Failure		400							{object}	fail.Response
//	@Failure		500							{object}	fail.Response
//	@Router			/v1/shopping-lists/personal	[get]
func (h *Handler) GetPersonalShoppingList(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetShoppingList(c, &api.GetShoppingListRequest{
		UserId: payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetShoppingList(res))
}

// GetShoppingList Swagger Documentation
//
//	@Summary		Get shopping list
//	@Description	Get shopping list
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			shopping_list_id						path		string							true	"Shopping list ID"
//	@Param			input									body		request_body.GetShoppingList	true	"Key"
//	@Success		200										{object}	response_body.GetShoppingListBody
//	@Failure		400										{object}	fail.Response
//	@Failure		500										{object}	fail.Response
//	@Router			/v1/shopping-lists/{shopping_list_id}	[get]
func (h *Handler) GetShoppingList(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.GetShoppingList
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}
	key := ""
	if body.Key != nil {
		key = *body.Key
	}

	res, err := h.service.GetShoppingList(c, &api.GetShoppingListRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		UserId:         payload.UserId.String(),
		Key:            key,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetShoppingList(res))
}

// SetShoppingListName Swagger Documentation
//
//	@Summary		Set shopping list
//	@Description	Set shopping list
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			shopping_list_id							path		int								true	"Shopping list ID"
//	@Param			input										body		request_body.SetShoppingList	true	"Shopping list"
//	@Success		200											{object}	response_body.SetShoppingList
//	@Failure		400											{object}	fail.Response
//	@Failure		500											{object}	fail.Response
//	@Router			/v1/shopping-lists/{shopping_list_id}/name	[put]
func (h *Handler) SetShoppingListName(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SetShoppingListName
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}
	name := ""
	if body.Name != nil {
		name = *body.Name
	}

	res, err := h.service.SetShoppingListName(c, &api.SetShoppingListNameRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		UserId:         payload.UserId.String(),
		Name:           name,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// SetShoppingList Swagger Documentation
//
//	@Summary		Set shopping list
//	@Description	Set shopping list
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			shopping_list_id						path		int								true	"Shopping list ID"
//	@Param			input									body		request_body.SetShoppingList	true	"Shopping list"
//	@Success		200										{object}	response_body.SetShoppingList
//	@Failure		400										{object}	fail.Response
//	@Failure		500										{object}	fail.Response
//	@Router			/v1/shopping-lists/{shopping_list_id}	[put]
func (h *Handler) SetShoppingList(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SetShoppingList
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}
	var lastVersion int32 = 0
	if body.LastVersion != nil {
		lastVersion = *body.LastVersion
	}

	res, err := h.service.SetShoppingList(c, &api.SetShoppingListRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		EditorId:       payload.UserId.String(),
		Purchases:      request_body.Purchases(body.Purchases),
		RecipeNames:    body.RecipeNames,
		LastVersion:    lastVersion,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.SetShoppingList{Version: res.Version})
}

// AddToShoppingList Swagger Documentation
//
//	@Summary		Add to shopping list
//	@Description	Add new purchases to shopping list
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			shopping_list_id						path		int								true	"Shopping list ID"
//	@Param			input									body		request_body.SetShoppingList	true	"Purchases"
//	@Success		200										{object}	response_body.SetShoppingList
//	@Failure		400										{object}	fail.Response
//	@Failure		500										{object}	fail.Response
//	@Router			/v1/shopping-lists/{shopping_list_id}	[patch]
func (h *Handler) AddToShoppingList(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SetShoppingList
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}
	var lastVersion int32 = 0
	if body.LastVersion != nil {
		lastVersion = *body.LastVersion
	}

	res, err := h.service.AddPurchasesToShoppingList(c, &api.SetShoppingListRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		EditorId:       payload.UserId.String(),
		Purchases:      request_body.Purchases(body.Purchases),
		RecipeNames:    body.RecipeNames,
		LastVersion:    lastVersion,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.SetShoppingList{Version: res.Version})
}

// DeleteSharedShoppingList Swagger Documentation
//
//	@Summary		Delete shared shopping list
//	@Description	Delete shared shopping list
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			shopping_list_id						path		int	true	"Shopping list ID"
//	@Success		200										{object}	response.MessageBody
//	@Failure		400										{object}	fail.Response
//	@Failure		500										{object}	fail.Response
//	@Router			/v1/shopping-lists/{shopping_list_id}	[delete]
func (h *Handler) DeleteSharedShoppingList(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.DeleteSharedShoppingList(c, &api.DeleteSharedShoppingListRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		UserId:         payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
