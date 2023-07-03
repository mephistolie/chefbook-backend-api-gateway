package encryption

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/encryption/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/encryption/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-encryption/api/proto/implementation/v1"
)

// GetRecipeKeyRequests Swagger Documentation
//
//	@Summary		Get recipe key requests
//	@Description	Get recipe key access requests
//	@Tags			encryption, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id									path		string	true	"Recipe ID"
//	@Success		200											{object}	[]response_body.RecipeKeyRequest
//	@Failure		400											{object}	fail.Response
//	@Failure		500											{object}	fail.Response
//	@Router			/v1/encryption/recipes/{recipe_id}/users	[get]
func (h *Handler) GetRecipeKeyRequests(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetRecipeKeyRequests(c, &api.GetRecipeKeyRequestsRequest{
		RecipeId: c.Param(ParamRecipeId),
		UserId:   payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	requests := make([]response_body.RecipeKeyRequest, len(res.Requests))
	for i, keyRequest := range res.Requests {
		var keyPtr *string
		if len(keyRequest.PublicKey) > 0 {
			key := string(keyRequest.PublicKey[:])
			keyPtr = &key
		}

		requests[i] = response_body.RecipeKeyRequest{
			UserId:     keyRequest.UserId,
			UserName:   keyRequest.UserName,
			UserAvatar: keyRequest.UserAvatar,
			Status:     keyRequest.Status,
			PublicKey:  keyPtr,
		}
	}

	response.Success(c, requests)
}

// RequestRecipeKeyAccess Swagger Documentation
//
//	@Summary		Request recipe key access
//	@Description	Request recipe key access
//	@Tags			encryption, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id									path		string	true	"Recipe ID"
//	@Success		200											{object}	response.MessageBody
//	@Failure		400											{object}	fail.Response
//	@Failure		500											{object}	fail.Response
//	@Router			/v1/encryption/recipes/{recipe_id}/users	[post]
func (h *Handler) RequestRecipeKeyAccess(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.RequestRecipeKeyAccess(c, &api.RequestRecipeKeyAccessRequest{
		RecipeId: c.Param(ParamRecipeId),
		UserId:   payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// GetRecipeKey Swagger Documentation
//
//	@Summary		Get recipe key
//	@Description	Get recipe encrypted key
//	@Tags			encryption, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id							path		string	true	"Recipe ID"
//	@Success		200									{object}	[]response_body.GetRecipeKey
//	@Failure		400									{object}	fail.Response
//	@Failure		500									{object}	fail.Response
//	@Router			/v1/encryption/recipes/{recipe_id}	[get]
func (h *Handler) GetRecipeKey(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetRecipeKey(c, &api.GetRecipeKeyRequest{
		RecipeId: c.Param(ParamRecipeId),
		UserId:   payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}
	var keyPtr *string
	if len(res.EncryptedKey) > 0 {
		key := string(res.EncryptedKey[:])
		keyPtr = &key
	}

	response.Success(c, response_body.GetEncryptedVaultKey{Key: keyPtr})
}

// SetRecipeOwnerKey Swagger Documentation
//
//	@Summary		Set recipe owner key
//	@Description	Set recipe owner key
//	@Tags			encryption, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id							path		string						true	"Recipe ID"
//	@Param			input								body		request_body.SetRecipeKey	true	"Key"
//	@Success		200									{object}	response.MessageBody
//	@Failure		400									{object}	fail.Response
//	@Failure		500									{object}	fail.Response
//	@Router			/v1/encryption/recipes/{recipe_id}	[post]
func (h *Handler) SetRecipeOwnerKey(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SetRecipeKey
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.SetRecipeKey(c, &api.SetRecipeKeyRequest{
		RecipeId:     c.Param(ParamRecipeId),
		RequesterId:  payload.UserId.String(),
		EncryptedKey: []byte(body.Key),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// GrantRecipeKeyAccess Swagger Documentation
//
//	@Summary		Set recipe user key
//	@Description	Set recipe user key
//	@Tags			encryption, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id											path		string						true	"Recipe ID"
//	@Param			user_id												path		string						true	"User ID"
//	@Param			input												body		request_body.SetRecipeKey	true	"Key"
//	@Success		200													{object}	response.MessageBody
//	@Failure		400													{object}	fail.Response
//	@Failure		500													{object}	fail.Response
//	@Router			/v1/encryption/recipes/{recipe_id}/users/{user_id}	[post]
func (h *Handler) GrantRecipeKeyAccess(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SetRecipeKey
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	userId := c.Param(ParamUserId)

	res, err := h.service.SetRecipeKey(c, &api.SetRecipeKeyRequest{
		RecipeId:     c.Param(ParamRecipeId),
		UserId:       &userId,
		RequesterId:  payload.UserId.String(),
		EncryptedKey: []byte(body.Key),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// DeclineRecipeKeyAccess Swagger Documentation
//
//	@Summary		Delete recipe key
//	@Description	Delete recipe key
//	@Tags			encryption, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id											path		string	true	"Recipe ID"
//	@Param			user_id												path		string	true	"User ID"
//	@Success		200													{object}	response.MessageBody
//	@Failure		400													{object}	fail.Response
//	@Failure		500													{object}	fail.Response
//	@Router			/v1/encryption/recipes/{recipe_id}/users/{user_id}	[delete]
func (h *Handler) DeclineRecipeKeyAccess(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.DeleteRecipeKey(c, &api.DeleteRecipeKeyRequest{
		RecipeId:    c.Param(ParamRecipeId),
		UserId:      c.Param(ParamUserId),
		RequesterId: payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
