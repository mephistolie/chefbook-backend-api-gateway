package v1

import (
	"github.com/gin-gonic/gin"
)

func (r *Router) initRecipesRoutes(api *gin.RouterGroup) {
	recipesGroup := api.Group("/recipes", r.authMiddleware.AuthorizeUser)
	r.initCategoriesRoutes(recipesGroup)
	r.initTagsRoutes(recipesGroup)
}
