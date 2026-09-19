package shopping_list

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
)

func (h *Handler) GetShoppingListUsers(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetShoppingListUsers(c, &api.GetShoppingListUsersRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		RequesterId:    payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, gin.H{"users": response_body.ShoppingListUsers(res.Users)})
}

func (h *Handler) GetSharedShoppingListLink(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetSharedShoppingListLink(c, &api.GetSharedShoppingListLinkRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		RequesterId:    payload.UserId.String(),
		LinkPattern:    h.routes.JoinShoppingList,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetShoppingListLink{Link: res.Link, ExpirationTimestamp: res.ExpiresAt.AsTime()})
}

func (h *Handler) JoinShoppingList(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.JoinShoppingList
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.JoinShoppingList(c, &api.JoinShoppingListRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		UserId:         payload.UserId.String(),
		Key:            body.Key,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

func (h *Handler) DeleteUserFromShoppingList(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.DeleteUserFromShoppingList(c, &api.DeleteUserFromShoppingListRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		UserId:         c.Param(ParamUserId),
		RequesterId:    payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
