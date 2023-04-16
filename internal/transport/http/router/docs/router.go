package docs

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mephistolie/chefbook-backend-api-gateway/docs"
	"github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init(api *gin.RouterGroup) {
	api.GET("/doc/*any", swagger.WrapHandler(swaggerFiles.Handler))
}
