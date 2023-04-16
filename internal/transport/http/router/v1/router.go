package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/middleware/auth"
)

type Router struct {
	handler        v1.Handler
	authMiddleware *auth.Middleware
}

func NewRouter(handler v1.Handler, authMiddleware *auth.Middleware) *Router {
	return &Router{
		handler:        handler,
		authMiddleware: authMiddleware,
	}
}

func (r *Router) Init(api *gin.RouterGroup) {
	routerGroup := api.Group("/v1")
	{
		r.initAuthRoutes(routerGroup)
		r.initProfileRoutes(routerGroup)
	}
}

func (r *Router) initAuthRoutes(api *gin.RouterGroup) {
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/sign-up", r.handler.Auth.SignUp)
		authGroup.GET("/activate", r.handler.Auth.ActivateProfile)
		authGroup.POST("/sign-in", r.handler.Auth.SignIn)
		authGroup.POST("/refresh", r.handler.Auth.RefreshSession)
		authGroup.POST("/sign-out", r.handler.Auth.SignOut)

		authGroup.GET("/google/request", r.handler.Auth.RequestGoogleOAuth)
		authGroup.POST("/google", r.handler.Auth.SignInGoogle)
		authGroup.PUT("/google", r.authMiddleware.AuthorizeUser, r.handler.Auth.ConnectGoogle)
		authGroup.DELETE("/google", r.authMiddleware.AuthorizeUser, r.handler.Auth.DeleteGoogleConnection)

		authGroup.GET("/vk/request", r.handler.Auth.RequestVkOAuth)
		authGroup.POST("/vk", r.handler.Auth.SignInVk)
		authGroup.PUT("/vk", r.authMiddleware.AuthorizeUser, r.handler.Auth.ConnectVk)
		authGroup.DELETE("/vk", r.authMiddleware.AuthorizeUser, r.handler.Auth.DeleteVkConnection)

		authGroup.GET("/sessions", r.authMiddleware.AuthorizeUser, r.handler.Auth.GetSessions)
		authGroup.DELETE("/sessions", r.authMiddleware.AuthorizeUser, r.handler.Auth.EndSessions)

		authGroup.GET("/password", r.handler.Auth.RequestPasswordReset)
		authGroup.POST("/password", r.handler.Auth.ResetPassword)
		authGroup.PUT("/password", r.authMiddleware.AuthorizeUser, r.handler.Auth.ChangePassword)

		authGroup.GET("/nickname", r.authMiddleware.AuthorizeUser, r.handler.Auth.CheckNicknameAvailability)
		authGroup.POST("/nickname", r.authMiddleware.AuthorizeUser, r.handler.Auth.SetNickname)
	}
}

func (r *Router) initProfileRoutes(api *gin.RouterGroup) {
	authGroup := api.Group("/profile", r.authMiddleware.AuthorizeUser)
	{
		authGroup.DELETE("", r.handler.Profile.DeleteProfile)
	}
}
