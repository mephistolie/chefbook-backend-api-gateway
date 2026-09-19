package handler

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	v1 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1"
)

type Handler struct {
	V1 v1.Handler
}

func NewHandler(services *service.Services, cfg *config.Config) *Handler {
	return &Handler{
		V1: *v1.NewHandler(services, cfg),
	}
}
