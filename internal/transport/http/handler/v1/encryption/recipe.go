package encryption

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/encryption/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/encryption/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-encryption/api/proto/implementation/v1"
)

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
			key := base64.StdEncoding.EncodeToString(keyRequest.PublicKey)
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

	response.Success(c, gin.H{"requests": requests})
}

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
		key := base64.StdEncoding.EncodeToString(res.EncryptedKey)
		keyPtr = &key
	}

	response.Success(c, response_body.GetEncryptedVaultKey{Key: keyPtr})
}

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
	key, err := base64.StdEncoding.DecodeString(body.Key)
	if err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.SetRecipeKey(c, &api.SetRecipeKeyRequest{
		RecipeId:     c.Param(ParamRecipeId),
		RequesterId:  payload.UserId.String(),
		EncryptedKey: key,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

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
	key, err := base64.StdEncoding.DecodeString(body.Key)
	if err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	userId := c.Param(ParamUserId)

	res, err := h.service.SetRecipeKey(c, &api.SetRecipeKeyRequest{
		RecipeId:     c.Param(ParamRecipeId),
		UserId:       &userId,
		RequesterId:  payload.UserId.String(),
		EncryptedKey: key,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

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
