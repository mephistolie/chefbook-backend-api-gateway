package recipe

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
)

const (
	ParamRecipeId     = "recipe_id"
	ParamLanguageCode = "language_code"
)

type Handler struct {
	service *service.Recipe
}

func NewHandler(service *service.Recipe) *Handler {
	return &Handler{service: service}
}
