package profile

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
)

const (
	ParamProfileId = "profile_id"
)

type Handler struct {
	auth    *service.Auth
	user    *service.User
	profile *service.Profile
}

func NewHandler(auth *service.Auth, user *service.User, profile *service.Profile) *Handler {
	return &Handler{
		auth:    auth,
		user:    user,
		profile: profile,
	}
}
