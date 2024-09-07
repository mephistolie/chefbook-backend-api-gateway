package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe"
)

func (r *Router) initCollectionsRoutes(api *gin.RouterGroup) {
	collectionsGroup := api.Group("/collections", r.authMiddleware.AuthorizeUser)
	{
		collectionsGroup.GET("", r.handler.Recipe.GetCollections)
		collectionsGroup.POST("", r.handler.Recipe.AddCollection)
		collectionsGroup.GET(fmt.Sprintf("/:%s", recipe.ParamCollectionId), r.handler.Recipe.GetCollection)
		collectionsGroup.PUT(fmt.Sprintf("/:%s", recipe.ParamCollectionId), r.handler.Recipe.UpdateCollection)
		collectionsGroup.DELETE(fmt.Sprintf("/:%s", recipe.ParamCollectionId), r.handler.Recipe.DeleteCollection)

		collectionsGroup.POST(fmt.Sprintf("/:%s/save", recipe.ParamCollectionId), r.handler.Recipe.SaveCollectionToRecipeBook)
		collectionsGroup.DELETE(fmt.Sprintf("/:%s/save", recipe.ParamCollectionId), r.handler.Recipe.RemoveCollectionFromRecipeBook)
	}
}
