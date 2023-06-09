package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile"
)

func (r *Router) initProfileRoutes(api *gin.RouterGroup) {
	profileGroup := api.Group("/profile", r.authMiddleware.AuthorizeUser)
	{
		profileGroup.GET("", r.handler.Profile.GetProfile)
		profileGroup.PUT("/name", r.handler.Profile.SetName)
		profileGroup.PUT("/description", r.handler.Profile.SetDescription)
		profileGroup.POST("/avatar", r.handler.Profile.GenerateAvatarUploadLink)
		profileGroup.PUT("/avatar", r.handler.Profile.ConfirmAvatarUploading)
		profileGroup.DELETE("/avatar", r.handler.Profile.DeleteAvatar)
	}
	{
		api.GET("profile/delete", r.authMiddleware.AuthorizeDeletedUser, r.handler.Profile.GetProfileDeletionStatus)
		api.POST("profile/delete", r.authMiddleware.AuthorizeDeletedUser, r.handler.Profile.DeleteProfile)
		api.DELETE("profile/delete", r.authMiddleware.AuthorizeDeletedUser, r.handler.Profile.CancelProfileDeletion)
	}

	profilesGroup := api.Group("/profiles", r.authMiddleware.AuthorizeUser)
	{
		profilesGroup.GET(fmt.Sprintf("/:%s", profile.ParamProfileId), r.handler.Profile.GetProfile)
	}
}
