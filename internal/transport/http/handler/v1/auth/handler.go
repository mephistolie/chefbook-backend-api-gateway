package auth

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
)

type Handler struct{ service *service.Auth }

// Callback allowlists and mail links belong to the auth service configuration.
func NewHandler(service *service.Auth, _ config.Domains) *Handler { return &Handler{service: service} }
