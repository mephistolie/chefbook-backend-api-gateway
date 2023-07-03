package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/category"
)

func (r *Router) initCategoriesRoutes(api *gin.RouterGroup) {
	categoriesGroup := api.Group("/categories")
	{
		categoriesGroup.GET("", r.handler.Category.GetCategories)
		categoriesGroup.POST("", r.handler.Category.AddCategory)
		categoriesGroup.GET(fmt.Sprintf("/:%s", category.ParamCategoryId), r.handler.Category.GetCategory)
		categoriesGroup.PUT(fmt.Sprintf("/:%s", category.ParamCategoryId), r.handler.Category.UpdateCategory)
		categoriesGroup.DELETE(fmt.Sprintf("/:%s", category.ParamCategoryId), r.handler.Category.DeleteCategory)
	}
}
