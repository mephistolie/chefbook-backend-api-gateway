package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/middleware/auth"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/middleware/log"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/router/docs"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/router/v1"
	"github.com/mephistolie/chefbook-backend-api-gateway/pkg/limiter"
)

type Router struct {
	v1   v1.Router
	docs docs.Router
}

func NewRouter(handler *handler.Handler, authMiddleware *auth.Middleware) *Router {
	return &Router{
		v1:   *v1.NewRouter(handler.V1, authMiddleware),
		docs: *docs.NewRouter(),
	}
}

func (r *Router) Init(cfg *config.Config) *gin.Engine {
	mode := gin.DebugMode
	if *cfg.Environment == config.EnvProd {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	engine := gin.New()

	engine.Use(
		gin.Recovery(),
		log.Middleware(),
		limiter.Limit(*cfg.Limiter.RPS, *cfg.Limiter.Burst, *cfg.Limiter.TTL),
	)

	r.initAPI(engine)

	return engine
}

func (r *Router) initAPI(router *gin.Engine) {
	api := router.Group("/")
	{
		r.v1.Init(api)

		if gin.Mode() != gin.ReleaseMode {
			r.docs.Init(api)
		}
	}
}
