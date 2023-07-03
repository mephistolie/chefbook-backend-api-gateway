package encryption

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/encryption/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/encryption/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-encryption/api/proto/implementation/v1"
)

// GetEncryptedVaultKey Swagger Documentation
//
//	@Summary		Get encrypted vault key
//	@Description	Get profile  encrypted vault key
//	@Tags			encryption, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200						{object}	[]response_body.GetEncryptedVaultKey
//	@Failure		400						{object}	fail.Response
//	@Failure		500						{object}	fail.Response
//	@Router			/v1/encryption/vault	[get]
func (h *Handler) GetEncryptedVaultKey(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetEncryptedVaultKey(c, &api.GetEncryptedVaultKeyRequest{UserId: payload.UserId.String()})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}
	var keyPtr *string
	if len(res.EncryptedPrivateKey) > 0 {
		key := string(res.EncryptedPrivateKey[:])
		keyPtr = &key
	}

	response.Success(c, response_body.GetEncryptedVaultKey{Key: keyPtr})
}

// CreateEncryptedVault Swagger Documentation
//
//	@Summary		Create encrypted vault
//	@Description	Create profile encrypted vault
//	@Tags			encryption, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input					body		request_body.CreateEncryptedVault	true	"Keys"
//	@Success		200						{object}	response.MessageBody
//	@Failure		400						{object}	fail.Response
//	@Failure		500						{object}	fail.Response
//	@Router			/v1/encryption/vault	[post]
func (h *Handler) CreateEncryptedVault(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.CreateEncryptedVault
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.CreateEncryptedVault(c, &api.CreateEncryptedVaultRequest{
		UserId:              payload.UserId.String(),
		PublicKey:           []byte(body.PublicKey),
		EncryptedPrivateKey: []byte(body.PrivateKey),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// RequestEncryptedVaultDeletion Swagger Documentation
//
//	@Summary		Request encrypted vault deletion
//	@Description	Request profile encrypted vault deletion
//	@Tags			encryption, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200							{object}	response.MessageBody
//	@Failure		400							{object}	fail.Response
//	@Failure		500							{object}	fail.Response
//	@Router			/v1/encryption/vault/delete	[post]
func (h *Handler) RequestEncryptedVaultDeletion(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.RequestEncryptedVaultDeletion(c, &api.RequestEncryptedVaultDeletionRequest{
		UserId: payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// DeleteEncryptedVault Swagger Documentation
//
//	@Summary		Delete encrypted vault
//	@Description	Delete profile encrypted vault
//	@Tags			encryption, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input					body		request_body.CreateEncryptedVault	true	"Keys"
//	@Success		200						{object}	response.MessageBody
//	@Failure		400						{object}	fail.Response
//	@Failure		500						{object}	fail.Response
//	@Router			/v1/encryption/vault	[delete]
func (h *Handler) DeleteEncryptedVault(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.DeleteEncryptedVault
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.DeleteEncryptedVault(c, &api.DeleteEncryptedVaultRequest{
		UserId:     payload.UserId.String(),
		DeleteCode: body.DeleteCode,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
