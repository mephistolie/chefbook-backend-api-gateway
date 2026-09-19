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
		key := base64.StdEncoding.EncodeToString(res.EncryptedPrivateKey)
		keyPtr = &key
	}
	var saltPtr *string
	if len(res.Salt) > 0 {
		salt := base64.StdEncoding.EncodeToString(res.Salt)
		saltPtr = &salt
	}

	response.Success(c, response_body.GetEncryptedVaultKey{Key: keyPtr, Salt: saltPtr})
}

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

	publicKey, err := base64.StdEncoding.DecodeString(body.PublicKey)
	if err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}
	privateKey, err := base64.StdEncoding.DecodeString(body.PrivateKey)
	if err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}
	salt, err := base64.StdEncoding.DecodeString(body.Salt)
	if err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.CreateEncryptedVault(c, &api.CreateEncryptedVaultRequest{
		UserId:              payload.UserId.String(),
		PublicKey:           publicKey,
		EncryptedPrivateKey: privateKey,
		Salt:                salt,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

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
