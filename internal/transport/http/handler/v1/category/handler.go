package category

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
)

const (
	ParamCategoryId = "category_id"
)

type Handler struct {
	service *service.Category
}

func NewHandler(service *service.Category) *Handler {
	return &Handler{service: service}
}
