package subscription

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
)

type Handler struct {
	service *service.Subscription
}

func NewHandler(service *service.Subscription) *Handler {
	return &Handler{service: service}
}
