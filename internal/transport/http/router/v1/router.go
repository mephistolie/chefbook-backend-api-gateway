package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list"

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
		r.initShoppingListRoutes(routerGroup)
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
	profileGroup := api.Group("/profile", r.authMiddleware.AuthorizeUser)
	{
		profileGroup.DELETE("", r.handler.Profile.DeleteProfile)
	}
}

func (r *Router) initShoppingListRoutes(api *gin.RouterGroup) {
	shoppingListGroup := api.Group("/shopping-lists", r.authMiddleware.AuthorizeUser)
	{
		shoppingListGroup.GET("", r.handler.ShoppingList.GetShoppingLists)

		shoppingListGroup.POST("", r.handler.ShoppingList.CreateSharedShoppingList)
		shoppingListGroup.GET("/personal", r.handler.ShoppingList.GetPersonalShoppingList)
		shoppingListGroup.GET(fmt.Sprintf("/:%s", shopping_list.ParamShoppingListId), r.handler.ShoppingList.GetShoppingList)
		shoppingListGroup.PUT(fmt.Sprintf("/:%s/name", shopping_list.ParamShoppingListId), r.handler.ShoppingList.SetShoppingListName)
		shoppingListGroup.PUT(fmt.Sprintf("/:%s", shopping_list.ParamShoppingListId), r.handler.ShoppingList.SetShoppingList)
		shoppingListGroup.PATCH(fmt.Sprintf("/:%s", shopping_list.ParamShoppingListId), r.handler.ShoppingList.AddToShoppingList)
		shoppingListGroup.DELETE(fmt.Sprintf("/:%s", shopping_list.ParamShoppingListId), r.handler.ShoppingList.DeleteSharedShoppingList)

		shoppingListGroup.GET(fmt.Sprintf("/:%s/link", shopping_list.ParamShoppingListId), r.handler.ShoppingList.GetSharedShoppingListLink)
		shoppingListGroup.GET(fmt.Sprintf("/:%s/users", shopping_list.ParamShoppingListId), r.handler.ShoppingList.GetShoppingListUsers)
		shoppingListGroup.POST(fmt.Sprintf("/:%s/users", shopping_list.ParamShoppingListId), r.handler.ShoppingList.JoinShoppingList)
		shoppingListGroup.DELETE(fmt.Sprintf("/:%s/users", shopping_list.ParamShoppingListId), r.handler.ShoppingList.DeleteUserFromShoppingList)
	}
}
