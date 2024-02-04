package auth

import (
	"fmt"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
)

const (
	ParamNickname = "nickname"
)

type Handler struct {
	service *service.Auth
	routes  Routes
}

type Routes struct {
	ActivateProfile string
	ResetPassword   string
	SignInGoogle    string
	SignInVk        string
}

func NewHandler(service *service.Auth, cfg config.Domains) *Handler {
	return &Handler{
		service: service,
		routes: Routes{
			ActivateProfile: fmt.Sprint("https://", *cfg.Frontend, "/auth/activate?user_id=%s&code=%s"),
			ResetPassword:   fmt.Sprint("https://", *cfg.Frontend, "/auth/reset-password?user_id=%s&code=%s"),
			SignInGoogle:    fmt.Sprintf("https://%s/auth/google", *cfg.Frontend),
			SignInVk:        fmt.Sprintf("https://%s/auth/vk", *cfg.Frontend),
		},
	}
}
