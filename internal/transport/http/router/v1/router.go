package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/middleware/auth"
)

type Router struct {
	handler        v1.Handler
	authMiddleware *auth.Middleware
}

func NewRouter(handler v1.Handler, authMiddleware *auth.Middleware) *Router {
	return &Router{
		handler:        handler,
		authMiddleware: authMiddleware,
	}
}

func (r *Router) Init(api *gin.RouterGroup) {
	routerGroup := api.Group("/v1")
	{
		r.initAuthRoutes(routerGroup)
		r.initSubscriptionRoutes(routerGroup)
		r.initProfileRoutes(routerGroup)
		r.initRecipesRoutes(routerGroup)
		r.initEncryptionRoutes(routerGroup)
		r.initShoppingListsRoutes(routerGroup)
	}
}
