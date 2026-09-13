package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	eventlog "github.com/mephistolie/chefbook-backend-api-gateway/internal/logging"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	"github.com/mephistolie/chefbook-backend-common/tokens/access"
)

const (
	userPayloadKey = "userPayload"
)

func PutUserPayload(c *gin.Context, payload access.Payload) {
	c.Set(userPayloadKey, payload)
}

func GetUserPayloadOrResponse(c *gin.Context) (*access.Payload, error) {
	payload, err := getUserPayload(c)
	if err != nil {
		eventlog.NewEvents().UserPayloadReadFailed(c.Request.Context(), err)
		response.Unknown(c, err)
	}
	return payload, err
}

func getUserPayload(c *gin.Context) (*access.Payload, error) {
	rawPayload, ok := c.Get(userPayloadKey)
	if !ok {
		return nil, errors.New("user payload not found")
	}

	payload, ok := rawPayload.(access.Payload)
	if !ok {
		return nil, errors.New("failed to parse user payload")
	}

	return &payload, nil
}
