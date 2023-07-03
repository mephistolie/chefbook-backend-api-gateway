package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe"
)

func (r *Router) initRecipesRoutes(api *gin.RouterGroup) {
	recipesGroup := api.Group("/recipes", r.authMiddleware.AuthorizeUser)
	r.initBaseRecipesRoutes(recipesGroup)
	r.initCategoriesRoutes(recipesGroup)
	r.initTagsRoutes(recipesGroup)
}

func (r *Router) initBaseRecipesRoutes(recipesGroup *gin.RouterGroup) {
	recipesGroup.GET("", r.handler.Recipe.GetRecipes)
	recipesGroup.GET("/random", r.handler.Recipe.GetRandomRecipe)
	recipesGroup.GET("/book", r.handler.Recipe.GetRecipeBook)

	recipesGroup.POST("", r.handler.Recipe.CreateRecipe)
	recipesGroup.GET(fmt.Sprintf("/:%s", recipe.ParamRecipeId), r.handler.Recipe.GetRecipe)
	recipesGroup.PUT(fmt.Sprintf("/:%s", recipe.ParamRecipeId), r.handler.Recipe.UpdateRecipe)
	recipesGroup.DELETE(fmt.Sprintf("/:%s", recipe.ParamRecipeId), r.handler.Recipe.DeleteRecipe)

	recipesGroup.POST(fmt.Sprintf("/:%s/rate", recipe.ParamRecipeId), r.handler.Recipe.RateRecipe)
	recipesGroup.POST(fmt.Sprintf("/:%s/save", recipe.ParamRecipeId), r.handler.Recipe.SaveRecipe)
	recipesGroup.DELETE(fmt.Sprintf("/:%s/save", recipe.ParamRecipeId), r.handler.Recipe.RemoveRecipeFromRecipeBook)
	recipesGroup.POST(fmt.Sprintf("/:%s/favourite", recipe.ParamRecipeId), r.handler.Recipe.AddRecipeToFavourite)
	recipesGroup.DELETE(fmt.Sprintf("/:%s/favourite", recipe.ParamRecipeId), r.handler.Recipe.RemoveRecipeFromFavourite)
	recipesGroup.PUT(fmt.Sprintf("/:%s/categories", recipe.ParamRecipeId), r.handler.Recipe.SetRecipeCategories)
}
