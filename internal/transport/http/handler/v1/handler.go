package v1

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile"
)

type Handler struct {
	Auth    *auth.Handler
	Profile *profile.Handler
}

func NewHandler(services service.Services, cfg *config.Config) *Handler {
	return &Handler{
		Auth:    auth.NewHandler(services.Auth, cfg.Domains),
		Profile: profile.NewHandler(services.Auth),
	}
}
