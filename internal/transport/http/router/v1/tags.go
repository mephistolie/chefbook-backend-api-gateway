package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tag "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/tag"
)

func (r *Router) initTagsRoutes(api *gin.RouterGroup) {
	tagsGroup := api.Group("/tags")
	{
		tagsGroup.GET("", r.handler.Tag.GetTags)
		tagsGroup.GET(fmt.Sprintf("/:%s", tag.ParamTagId), r.handler.Tag.GetTag)
		tagsGroup.GET("/groups", r.handler.Tag.GetTagGroups)
	}
}
