package shopping_list

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
	"net/url"
)

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

	response.Success(c, gin.H{"shoppingLists": response_body.GetShoppingLists(res)})
}

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
	name := ""
	if body.Name != nil {
		name = *body.Name
	}

	res, err := h.service.CreateSharedShoppingList(c, &api.CreateSharedShoppingListRequest{
		ShoppingListId:   body.ShoppingListId,
		Name:             name,
		UserId:           payload.UserId.String(),
		SubscriptionPlan: payload.SubscriptionPlan,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Created(c, "/v1/shopping-lists/"+url.PathEscape(res.ShoppingListId), response_body.CreateShoppingList{Id: res.ShoppingListId})
}

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

func (h *Handler) GetShoppingList(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetShoppingList(c, &api.GetShoppingListRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		UserId:         payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetShoppingList(res))
}

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

	res, err := h.service.SetShoppingListName(c, &api.SetShoppingListNameRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		UserId:         payload.UserId.String(),
		Name:           body.Name,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

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
		LastVersion:    lastVersion,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.SetShoppingList{Version: res.Version})
}

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
		LastVersion:    lastVersion,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.SetShoppingList{Version: res.Version})
}

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
