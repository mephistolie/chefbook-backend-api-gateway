package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe"
)

func (r *Router) initRecipesRoutes(api *gin.RouterGroup) {
	recipesGroup := api.Group("/recipes", r.authMiddleware.AuthorizeUser)
	r.initBaseRecipesRoutes(recipesGroup)
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

	recipesGroup.POST(fmt.Sprintf("/:%s/pictures", recipe.ParamRecipeId), r.handler.Recipe.GenerateRecipePicturesUploadLinks)
	recipesGroup.PUT(fmt.Sprintf("/:%s/pictures", recipe.ParamRecipeId), r.handler.Recipe.SetRecipePictures)

	recipesGroup.POST(fmt.Sprintf("/:%s/rate", recipe.ParamRecipeId), r.handler.Recipe.RateRecipe)
	recipesGroup.POST(fmt.Sprintf("/:%s/book", recipe.ParamRecipeId), r.handler.Recipe.SaveRecipeToRecipeBook)
	recipesGroup.DELETE(fmt.Sprintf("/:%s/book", recipe.ParamRecipeId), r.handler.Recipe.RemoveRecipeFromRecipeBook)
	recipesGroup.POST(fmt.Sprintf("/:%s/favourites", recipe.ParamRecipeId), r.handler.Recipe.SaveRecipeToFavourites)
	recipesGroup.DELETE(fmt.Sprintf("/:%s/favourites", recipe.ParamRecipeId), r.handler.Recipe.RemoveRecipeFromFavourites)
	recipesGroup.POST(fmt.Sprintf("/:%s/collections/:%s", recipe.ParamRecipeId, recipe.ParamCollectionId), r.handler.Recipe.AddRecipeToCollection)
	recipesGroup.DELETE(fmt.Sprintf("/:%s/collections/:%s", recipe.ParamRecipeId, recipe.ParamCollectionId), r.handler.Recipe.RemoveRecipeFromCollection)
	recipesGroup.PUT(fmt.Sprintf("/:%s/collections", recipe.ParamRecipeId), r.handler.Recipe.SetRecipeCollections)

	recipesGroup.POST(fmt.Sprintf("/:%s/translations", recipe.ParamRecipeId), r.handler.Recipe.TranslateRecipe)
	recipesGroup.DELETE(fmt.Sprintf("/:%s/translations/:%s", recipe.ParamRecipeId, recipe.ParamLanguageCode), r.handler.Recipe.DeleteRecipeTranslation)
}
