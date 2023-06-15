package v1

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/category"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list"
	tag "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/tag"
)

type Handler struct {
	Auth         *auth.Handler
	Profile      *profile.Handler
	Tag          *tag.Handler
	Category     *category.Handler
	ShoppingList *shopping_list.Handler
}

func NewHandler(services *service.Services, cfg *config.Config) *Handler {
	return &Handler{
		Auth:         auth.NewHandler(services.Auth, cfg.Domains),
		Profile:      profile.NewHandler(services.Auth, services.User, services.Profile),
		Tag:          tag.NewHandler(services.Tag),
		Category:     category.NewHandler(services.Category),
		ShoppingList: shopping_list.NewHandler(services.ShoppingList, cfg.Domains),
	}
}
