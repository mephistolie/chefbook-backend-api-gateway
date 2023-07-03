package encryption

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
)

const (
	ParamRecipeId = "recipe_id"
	ParamUserId   = "user_id"
)

type Handler struct {
	service *service.Encryption
}

func NewHandler(service *service.Encryption) *Handler {
	return &Handler{service: service}
}
