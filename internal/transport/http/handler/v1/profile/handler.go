package profile

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
)

type Handler struct {
	auth *service.Auth
	user *service.User
}

func NewHandler(auth *service.Auth, user *service.User) *Handler {
	return &Handler{
		auth: auth,
		user: user,
	}
}
