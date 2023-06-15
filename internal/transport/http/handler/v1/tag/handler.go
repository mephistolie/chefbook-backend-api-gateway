package category

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
)

const (
	ParamTagId = "tag_id"
)

type Handler struct {
	service *service.Tag
}

func NewHandler(service *service.Tag) *Handler {
	return &Handler{service: service}
}
