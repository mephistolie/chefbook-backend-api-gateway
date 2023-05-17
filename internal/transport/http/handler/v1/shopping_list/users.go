package shopping_list

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
)

// GetShoppingListUsers Swagger Documentation
//
//	@Summary		Get shopping list users
//	@Description	Get shopping list users
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			shopping_list_id							path		int	true	"Shopping list ID"
//	@Success		200											{object}	[]string
//	@Failure		400											{object}	fail.Response
//	@Failure		500											{object}	fail.Response
//	@Router			/v1/shopping-lists/{shopping_list_id}/users	[get]
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

	response.Success(c, res.Users)
}

// GetSharedShoppingListLink Swagger Documentation
//
//	@Summary		Get shared shopping list link
//	@Description	Get shared shopping list link
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			shopping_list_id							path		int	true	"Shopping list ID"
//	@Success		200											{object}	response_body.GetShoppingListLink
//	@Failure		400											{object}	fail.Response
//	@Failure		500											{object}	fail.Response
//	@Router			/v1/shopping-lists/{shopping_list_id}/link	[get]
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

	response.Success(c, response_body.GetShoppingListLink{Link: res.Link, ExpiresAt: res.ExpiresAt.AsTime()})
}

// JoinShoppingList Swagger Documentation
//
//	@Summary		Join shared shopping list
//	@Description	Join shared shopping list
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			shopping_list_id							path		int								true	"Shopping list ID"
//	@Param			input										body		request_body.JoinShoppingList	true	"Key"
//	@Success		200											{object}	response.MessageBody
//	@Failure		400											{object}	fail.Response
//	@Failure		500											{object}	fail.Response
//	@Router			/v1/shopping-lists/{shopping_list_id}/users	[post]
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

// DeleteUserFromShoppingList Swagger Documentation
//
//	@Summary		Delete user from shared shopping list
//	@Description	Delete user from shared shopping list
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			shopping_list_id							path		int										true	"Shopping list ID"
//	@Param			input										body		request_body.DeleteUserFromShoppingList	true	"User ID"
//	@Success		200											{object}	response.MessageBody
//	@Failure		400											{object}	fail.Response
//	@Failure		500											{object}	fail.Response
//	@Router			/v1/shopping-lists/{shopping_list_id}/users	[delete]
func (h *Handler) DeleteUserFromShoppingList(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.DeleteUserFromShoppingList
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.DeleteUserFromShoppingList(c, &api.DeleteUserFromShoppingListRequest{
		ShoppingListId: c.Param(ParamShoppingListId),
		UserId:         body.UserId.String(),
		RequesterId:    payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
