package v1

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/encryption"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/subscription"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/tag"
)

type Handler struct {
	Auth         *auth.Handler
	Subscription *subscription.Handler
	Profile      *profile.Handler
	Tag          *tag.Handler
	Recipe       *recipe.Handler
	Encryption   *encryption.Handler
	ShoppingList *shopping_list.Handler
}

func NewHandler(services *service.Services, cfg *config.Config) *Handler {
	return &Handler{
		Auth:         auth.NewHandler(services.Auth, cfg.Domains),
		Subscription: subscription.NewHandler(services.Subscription),
		Profile:      profile.NewHandler(services.Auth, services.User, services.Profile),
		Tag:          tag.NewHandler(services.Tag),
		Recipe:       recipe.NewHandler(services.Recipe),
		Encryption:   encryption.NewHandler(services.Encryption),
		ShoppingList: shopping_list.NewHandler(services.ShoppingList, cfg.Domains),
	}
}
