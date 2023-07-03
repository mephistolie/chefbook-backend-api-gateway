package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/encryption"
)

func (r *Router) initEncryptionRoutes(api *gin.RouterGroup) {
	encryptionGroup := api.Group("/encryption", r.authMiddleware.AuthorizeUser)
	{
		vaultGroup := encryptionGroup.Group("/vault")
		{
			vaultGroup.GET("", r.handler.Encryption.GetEncryptedVaultKey)
			vaultGroup.POST("", r.handler.Encryption.CreateEncryptedVault)
			vaultGroup.POST("/delete", r.handler.Encryption.RequestEncryptedVaultDeletion)
			vaultGroup.DELETE("", r.handler.Encryption.DeleteEncryptedVault)
		}
		recipesGroup := encryptionGroup.Group("/recipes")
		{
			recipeRoute := fmt.Sprintf("/:%s", encryption.ParamRecipeId)

			recipesGroup.GET(fmt.Sprintf("%s/users", recipeRoute), r.handler.Encryption.GetRecipeKeyRequests)
			recipesGroup.POST(fmt.Sprintf("%s/users", recipeRoute), r.handler.Encryption.RequestRecipeKeyAccess)
			recipesGroup.GET(recipeRoute, r.handler.Encryption.GetRecipeKey)
			recipesGroup.POST(recipeRoute, r.handler.Encryption.SetRecipeOwnerKey)
			recipesGroup.POST(fmt.Sprintf("%s/users/:%s", recipeRoute, encryption.ParamUserId), r.handler.Encryption.GrantRecipeKeyAccess)
			recipesGroup.DELETE(fmt.Sprintf("%s/users/:%s", recipeRoute, encryption.ParamUserId), r.handler.Encryption.DeclineRecipeKeyAccess)
		}
	}
}
