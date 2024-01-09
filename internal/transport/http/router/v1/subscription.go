package v1

import (
	"github.com/gin-gonic/gin"
)

func (r *Router) initSubscriptionRoutes(api *gin.RouterGroup) {
	authGroup := api.Group("/subscriptions", r.authMiddleware.AuthorizeUser)
	{
		authGroup.GET("", r.handler.Subscription.GetSubscriptions)
		authGroup.POST("/google", r.handler.Subscription.ConfirmGoogleSubscription)
	}
}
