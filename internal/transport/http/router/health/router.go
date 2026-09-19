package health

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init(api *gin.RouterGroup) {
	api.GET("/healthz", r.checkHealth)
}

func (r *Router) checkHealth(c *gin.Context) {
	response.Message(c, "API is alive")
}
