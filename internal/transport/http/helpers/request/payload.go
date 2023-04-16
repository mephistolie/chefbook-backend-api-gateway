package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-common/tokens/access"
)

const (
	userPayloadKey = "userPayload"
)

func PutUserPayload(c *gin.Context, payload access.Payload) {
	c.Set(userPayloadKey, payload)
}

func GetUserPayload(c *gin.Context) (*access.Payload, error) {
	rawPayload, ok := c.Get(userPayloadKey)
	if !ok {
		return nil, errors.New("unable to get user payload")
	}

	payload, ok := rawPayload.(access.Payload)
	if !ok {
		return nil, errors.New("unable to get user payload")
	}

	return &payload, nil
}
