package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/contract"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

var _ contract.ServerInterface = (*Router)(nil)

func (r *Router) GetCollections(c *gin.Context) {
	r.handler.Recipe.GetCollections(c)
}

func (r *Router) AddCollection(c *gin.Context) { r.handler.Recipe.AddCollection(c) }

func (r *Router) DeleteCollection(c *gin.Context, collectionId string) {
	r.handler.Recipe.DeleteCollection(c)
}

func (r *Router) GetCollection(c *gin.Context, collectionId string) {
	r.handler.Recipe.GetCollection(c)
}

func (r *Router) UpdateCollection(c *gin.Context, collectionId string) {
	r.handler.Recipe.UpdateCollection(c)
}

func (r *Router) RemoveCollectionFromRecipeBook(c *gin.Context, collectionId string) {
	r.handler.Recipe.RemoveCollectionFromRecipeBook(c)
}

func (r *Router) SaveCollectionToRecipeBook(c *gin.Context, collectionId string) {
	r.handler.Recipe.SaveCollectionToRecipeBook(c)
}

func (r *Router) GetRecipeKey(c *gin.Context, recipeId string) { r.handler.Encryption.GetRecipeKey(c) }

func (r *Router) SetRecipeOwnerKey(c *gin.Context, recipeId string) {
	r.handler.Encryption.SetRecipeOwnerKey(c)
}

func (r *Router) GetRecipeKeyRequests(c *gin.Context, recipeId string) {
	r.handler.Encryption.GetRecipeKeyRequests(c)
}

func (r *Router) RequestRecipeKeyAccess(c *gin.Context, recipeId string) {
	r.handler.Encryption.RequestRecipeKeyAccess(c)
}

func (r *Router) DeclineRecipeKeyAccess(c *gin.Context, recipeId string, userId string) {
	r.handler.Encryption.DeclineRecipeKeyAccess(c)
}

func (r *Router) GrantRecipeKeyAccess(c *gin.Context, recipeId string, userId string) {
	r.handler.Encryption.GrantRecipeKeyAccess(c)
}

func (r *Router) DeleteEncryptedVault(c *gin.Context) { r.handler.Encryption.DeleteEncryptedVault(c) }

func (r *Router) GetEncryptedVaultKey(c *gin.Context) { r.handler.Encryption.GetEncryptedVaultKey(c) }

func (r *Router) CreateEncryptedVault(c *gin.Context) { r.handler.Encryption.CreateEncryptedVault(c) }

func (r *Router) RequestEncryptedVaultDeletion(c *gin.Context) {
	r.handler.Encryption.RequestEncryptedVaultDeletion(c)
}

func (r *Router) GetProfile(c *gin.Context) { r.handler.Profile.GetProfile(c) }

func (r *Router) DeleteAvatar(c *gin.Context) { r.handler.Profile.DeleteAvatar(c) }

func (r *Router) GenerateAvatarUploadLink(c *gin.Context) {
	r.handler.Profile.GenerateAvatarUploadLink(c)
}

func (r *Router) ConfirmAvatarUploading(c *gin.Context) { r.handler.Profile.ConfirmAvatarUploading(c) }

func (r *Router) GetProfileDeletionStatus(c *gin.Context) {
	r.handler.Profile.GetProfileDeletionStatus(c)
}

func (r *Router) SetDescription(c *gin.Context) { r.handler.Profile.SetDescription(c) }

func (r *Router) SetDisplayName(c *gin.Context) { r.handler.Profile.SetDisplayName(c) }

func (r *Router) GetPublicProfile(c *gin.Context, profileId string) {
	r.handler.Profile.GetPublicProfile(c)
}

func (r *Router) GetRecipes(c *gin.Context, params contract.GetRecipesParams) {
	r.handler.Recipe.GetRecipes(c)
}

func (r *Router) CreateRecipe(c *gin.Context) { r.handler.Recipe.CreateRecipe(c) }

func (r *Router) GetRecipeBook(c *gin.Context, params contract.GetRecipeBookParams) {
	r.handler.Recipe.GetRecipeBook(c)
}

func (r *Router) GetRandomRecipe(c *gin.Context, params contract.GetRandomRecipeParams) {
	r.handler.Recipe.GetRandomRecipe(c)
}

func (r *Router) GetTags(c *gin.Context, params contract.GetTagsParams) { r.handler.Tag.GetTags(c) }

func (r *Router) GetTagGroups(c *gin.Context, params contract.GetTagGroupsParams) {
	r.handler.Tag.GetTagGroups(c)
}

func (r *Router) GetTag(c *gin.Context, tagId string, params contract.GetTagParams) {
	r.handler.Tag.GetTag(c)
}

func (r *Router) DeleteRecipe(c *gin.Context, recipeId string) { r.handler.Recipe.DeleteRecipe(c) }

func (r *Router) GetRecipe(c *gin.Context, recipeId string, params contract.GetRecipeParams) {
	r.handler.Recipe.GetRecipe(c)
}

func (r *Router) UpdateRecipe(c *gin.Context, recipeId string) { r.handler.Recipe.UpdateRecipe(c) }

func (r *Router) RemoveRecipeFromRecipeBook(c *gin.Context, recipeId string) {
	r.handler.Recipe.RemoveRecipeFromRecipeBook(c)
}

func (r *Router) SaveRecipeToRecipeBook(c *gin.Context, recipeId string) {
	r.handler.Recipe.SaveRecipeToRecipeBook(c)
}

func (r *Router) SetRecipeCollections(c *gin.Context, recipeId string) {
	r.handler.Recipe.SetRecipeCollections(c)
}

func (r *Router) RemoveRecipeFromCollection(c *gin.Context, recipeId string, collectionId string) {
	r.handler.Recipe.RemoveRecipeFromCollection(c)
}

func (r *Router) AddRecipeToCollection(c *gin.Context, recipeId string, collectionId string) {
	r.handler.Recipe.AddRecipeToCollection(c)
}

func (r *Router) RemoveRecipeFromFavourites(c *gin.Context, recipeId string) {
	r.handler.Recipe.RemoveRecipeFromFavourites(c)
}

func (r *Router) SaveRecipeToFavourites(c *gin.Context, recipeId string) {
	r.handler.Recipe.SaveRecipeToFavourites(c)
}

func (r *Router) GenerateRecipePicturesUploadLinks(c *gin.Context, recipeId string) {
	r.handler.Recipe.GenerateRecipePicturesUploadLinks(c)
}

func (r *Router) SetRecipePictures(c *gin.Context, recipeId string) {
	r.handler.Recipe.SetRecipePictures(c)
}

func (r *Router) RateRecipe(c *gin.Context, recipeId string) { r.handler.Recipe.RateRecipe(c) }

func (r *Router) TranslateRecipe(c *gin.Context, recipeId string) {
	r.handler.Recipe.TranslateRecipe(c)
}

func (r *Router) DeleteRecipeTranslation(c *gin.Context, recipeId string, languageCode string) {
	r.handler.Recipe.DeleteRecipeTranslation(c)
}

func (r *Router) GetShoppingLists(c *gin.Context) { r.handler.ShoppingList.GetShoppingLists(c) }

func (r *Router) CreateSharedShoppingList(c *gin.Context) {
	r.handler.ShoppingList.CreateSharedShoppingList(c)
}

func (r *Router) GetPersonalShoppingList(c *gin.Context) {
	r.handler.ShoppingList.GetPersonalShoppingList(c)
}

func (r *Router) DeleteSharedShoppingList(c *gin.Context, shoppingListId string) {
	r.handler.ShoppingList.DeleteSharedShoppingList(c)
}

func (r *Router) GetShoppingList(c *gin.Context, shoppingListId string) {
	r.handler.ShoppingList.GetShoppingList(c)
}

func (r *Router) AddToShoppingList(c *gin.Context, shoppingListId string) {
	r.handler.ShoppingList.AddToShoppingList(c)
}

func (r *Router) SetShoppingList(c *gin.Context, shoppingListId string) {
	r.handler.ShoppingList.SetShoppingList(c)
}

func (r *Router) GetSharedShoppingListLink(c *gin.Context, shoppingListId string) {
	r.handler.ShoppingList.GetSharedShoppingListLink(c)
}

func (r *Router) SetShoppingListName(c *gin.Context, shoppingListId string) {
	r.handler.ShoppingList.SetShoppingListName(c)
}

func (r *Router) GetShoppingListUsers(c *gin.Context, shoppingListId string) {
	r.handler.ShoppingList.GetShoppingListUsers(c)
}

func (r *Router) JoinShoppingList(c *gin.Context, shoppingListId string) {
	r.handler.ShoppingList.JoinShoppingList(c)
}

func (r *Router) DeleteUserFromShoppingList(c *gin.Context, shoppingListId string, userId string) {
	r.handler.ShoppingList.DeleteUserFromShoppingList(c)
}

func (r *Router) GetSubscriptions(c *gin.Context) { r.handler.Subscription.GetSubscriptions(c) }

func (r *Router) ConfirmGoogleSubscription(c *gin.Context) {
	r.handler.Subscription.ConfirmGoogleSubscription(c)
}

func (r *Router) GetBackupCodesStatus(c *gin.Context) { r.handler.Auth.GetBackupCodes(c) }

func (r *Router) GenerateBackupCodes(c *gin.Context, params contract.GenerateBackupCodesParams) {
	r.handler.Auth.CreateBackupCodes(c)
}

func (r *Router) CancelAccountDeletion(c *gin.Context) { r.handler.Auth.CancelProfileDeletion(c) }

func (r *Router) UpdateAccountDeletion(c *gin.Context) { r.handler.Auth.UpdateAccountDeletion(c) }

func (r *Router) RequestAccountDeletion(c *gin.Context, params contract.RequestAccountDeletionParams) {
	r.handler.Auth.DeleteProfile(c)
}

func (r *Router) RequestEmailChange(c *gin.Context, params contract.RequestEmailChangeParams) {
	r.handler.Auth.RequestEmailVerification(c)
}

func (r *Router) ConfirmEmailChange(c *gin.Context) { r.handler.Auth.ConfirmEmailChange(c) }

func (r *Router) GetIdentities(c *gin.Context) { r.handler.Auth.GetIdentities(c) }

func (r *Router) UnlinkGoogleIdentity(c *gin.Context, params contract.UnlinkGoogleIdentityParams) {
	r.handler.Auth.DeleteGoogleConnection(c)
}

func (r *Router) LinkGoogleIdentity(c *gin.Context, params contract.LinkGoogleIdentityParams) {
	r.handler.Auth.ConnectGoogle(c)
}

func (r *Router) UnlinkVkIdentity(c *gin.Context, params contract.UnlinkVkIdentityParams) {
	r.handler.Auth.DeleteVkConnection(c)
}

func (r *Router) LinkVkIdentity(c *gin.Context, params contract.LinkVkIdentityParams) {
	r.handler.Auth.ConnectVk(c)
}

func (r *Router) GetPasskeys(c *gin.Context) { r.handler.Auth.GetPasskeys(c) }

func (r *Router) ConfirmPasskeyRegistration(c *gin.Context) { r.handler.Auth.CreatePasskey(c) }

func (r *Router) RequestPasskeyRegistration(c *gin.Context, params contract.RequestPasskeyRegistrationParams) {
	r.handler.Auth.CreatePasskeyRegistration(c)
}

func (r *Router) DeletePasskey(c *gin.Context, id openapi_types.UUID, params contract.DeletePasskeyParams) {
	r.handler.Auth.DeletePasskey(c)
}

func (r *Router) RenamePasskey(c *gin.Context, id openapi_types.UUID) {
	r.handler.Auth.RenamePasskey(c)
}

func (r *Router) SetPassword(c *gin.Context, params contract.SetPasswordParams) {
	r.handler.Auth.ChangePassword(c)
}

func (r *Router) RequestPasswordReset(c *gin.Context) { r.handler.Auth.RequestPasswordReset(c) }

func (r *Router) ConfirmPasswordReset(c *gin.Context) { r.handler.Auth.ResetPassword(c) }

func (r *Router) DeleteTotp(c *gin.Context, params contract.DeleteTotpParams) {
	r.handler.Auth.DeleteTotp(c)
}

func (r *Router) GetTotpStatus(c *gin.Context) { r.handler.Auth.GetTotp(c) }

func (r *Router) RequestTotpActivation(c *gin.Context, params contract.RequestTotpActivationParams) {
	r.handler.Auth.CreateTotp(c)
}

func (r *Router) ConfirmTotpActivation(c *gin.Context) { r.handler.Auth.ConfirmTotp(c) }

func (r *Router) SetUsername(c *gin.Context) { r.handler.Auth.SetUsername(c) }

func (r *Router) StartAuthentication(c *gin.Context) { r.handler.Auth.CreateAuthentication(c) }

func (r *Router) GetAuthentication(c *gin.Context, id openapi_types.UUID, params contract.GetAuthenticationParams) {
	r.handler.Auth.GetAuthentication(c)
}

func (r *Router) StartAuthenticationStep(c *gin.Context, id openapi_types.UUID, params contract.StartAuthenticationStepParams) {
	r.handler.Auth.StartAuthenticationStep(c)
}

func (r *Router) GetAuthenticationStep(c *gin.Context, id openapi_types.UUID, stepId openapi_types.UUID, params contract.GetAuthenticationStepParams) {
	r.handler.Auth.GetAuthenticationStep(c)
}

func (r *Router) CompleteAuthenticationStep(c *gin.Context, id openapi_types.UUID, stepId openapi_types.UUID, params contract.CompleteAuthenticationStepParams) {
	r.handler.Auth.CompleteAuthenticationStep(c)
}

func (r *Router) CreateGoogleAuthorizationUrl(c *gin.Context) { r.handler.Auth.RequestGoogleOAuth(c) }

func (r *Router) CreateVkAuthorizationUrl(c *gin.Context) { r.handler.Auth.RequestVkOAuth(c) }

func (r *Router) RevokeAllSessions(c *gin.Context) { r.handler.Auth.EndSessions(c) }

func (r *Router) GetSessions(c *gin.Context) { r.handler.Auth.GetSessions(c) }

func (r *Router) CreateSession(c *gin.Context) { r.handler.Auth.CreateSession(c) }

func (r *Router) RevokeSession(c *gin.Context, id int64) { r.handler.Auth.EndSession(c, id) }

func (r *Router) RefreshSessionTokens(c *gin.Context, id int64) { r.handler.Auth.RefreshSession(c) }

func (r *Router) CheckUsernameAvailability(c *gin.Context, username string) {
	r.handler.Auth.CheckUsernameAvailability(c)
}
