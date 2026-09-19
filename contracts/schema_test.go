package contracts_test

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/contracts"
	dto13 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/request_body"
	dto12 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/response_body"
	dto15 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/encryption/dto/request_body"
	dto14 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/encryption/dto/response_body"
	dto6 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile/dto/request_body"
	dto5 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile/dto/response_body"
	dto11 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	dto10 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/request_body"
	dto9 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/response_body"
	dto3 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/common_body"
	dto2 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/request_body"
	dto1 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list/dto/response_body"
	dto8 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/subscription/dto/request_body"
	dto7 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/subscription/dto/response_body"
	dto4 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/tag/dto/response_body"
	dto0 "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	fail "github.com/mephistolie/chefbook-backend-common/responses/fail"
	"go.yaml.in/yaml/v3"
	"reflect"
	"strings"
	"testing"
)

var providerDTOs = map[string]any{
	"SetPasswordRequest":                 dto13.ChangePassword{},
	"UsernameAvailabilityResponse":          dto12.CheckUsername{},
	"RefreshSessionTokensRequest":                  dto13.RefreshToken{},
	"RequestPasswordResetRequest":           dto13.RequestPasswordReset{},
	"ConfirmPasswordResetRequest":           dto13.ResetPassword{},
	"Session":                               dto12.Session{},
	"TokensResponse":                        dto12.Tokens{},
	"SetUsernameRequest":                    dto13.Username{},
	"CreateEncryptedVaultRequest":           dto15.CreateEncryptedVault{},
	"ConfirmEncryptedVaultDeletionRequest":  dto15.DeleteEncryptedVault{},
	"EncryptedKeyResponse":                  dto14.GetEncryptedVaultKey{},
	"RecipeKeyAccessRequest":                dto14.RecipeKeyRequest{},
	"SetRecipeKeyRequest":                   dto15.SetRecipeKey{},
	"ErrorResponse":                         fail.Response{},
	"LinkResponse":                          dto0.LinkBody{},
	"MessageResponse":                       dto0.MessageBody{},
	"ConfirmAvatarUploadRequest":            dto6.ConfirmAvatarUploading{},
	"AvatarUploadResponse":                  dto5.GenerateAvatarUploadLink{},
	"ProfileSummary":                        dto0.ProfileInfo{},
	"ProfileSummaryFields":                  dto0.ProfileMinInfo{},
	"ProfileIdentities":                     dto5.OAuth{},
	"ProfileResponse":                       dto5.Profile{},
	"SetProfileDescriptionRequest":          dto6.SetDescription{},
	"SetDisplayNameRequest":                 dto6.SetDisplayName{},
	"CreateCollectionRequest":               dto10.AddCollection{},
	"CollectionCreatedResponse":             dto9.AddCollection{},
	"CollectionSummary":                     dto9.CollectionInfo{},
	"Collection":                            dto9.Collection{},
	"Contributor":                           dto9.Contributor{},
	"CookingStep":                           dto11.CookingItem{},
	"RecipeCreatedResponse":                 dto9.CreateRecipe{},
	"CreateRecipePictureUploadLinksRequest": dto10.GenerateRecipePicturesUploadLinks{},
	"CollectionResponse":                    dto9.GetCollection{},
	"CollectionsResponse":                   dto9.GetCollections{},
	"RecipeBookResponse":                    dto9.GetRecipeBookResponse{},
	"RecipeResponse":                        dto9.GetRecipeResponse{},
	"RecipesResponse":                       dto9.GetRecipesResponse{},
	"Ingredient":                            dto11.IngredientItem{},
	"IngredientTranslation":                 dto10.IngredientTranslation{},
	"Macronutrients":                        dto11.Macronutrients{},
	"RateRecipeRequest":                     dto10.RateRecipe{},
	"RecipeRating":                          dto9.Rating{},
	"RecipeSummary":                         dto9.RecipeInfo{},
	"SaveRecipeRequest":                     dto10.RecipeInput{},
	"RecipePictureUpload":                   dto9.RecipePictureUpload{},
	"RecipePictures":                        dto11.RecipePictures{},
	"Recipe":                                dto9.Recipe{},
	"RecipeState":                           dto9.RecipeState{},
	"SaveCollectionRequest":                 dto10.SaveCollectionToRecipeBook{},
	"SetRecipeCollectionsRequest":           dto10.SetRecipeCollections{},
	"SetRecipePicturesRequest":              dto10.SetRecipePictures{},
	"RecipePicturesResponse":                dto9.SetRecipePictures{},
	"RecipeTag":                             dto9.Tag{},
	"TranslateRecipeRequest":                dto10.TranslateRecipe{},
	"UpdateCollectionRequest":               dto10.UpdateCollection{},
	"RecipeVersionResponse":                 dto9.UpdateRecipe{},
	"CreateShoppingListRequest":             dto2.CreateSharedShoppingList{},
	"ShoppingListCreatedResponse":           dto1.CreateShoppingList{},
	"ShoppingListResponse":                  dto1.GetShoppingListBody{},
	"ShoppingListLinkResponse":              dto1.GetShoppingListLink{},
	"JoinShoppingListRequest":               dto2.JoinShoppingList{},
	"Purchase":                              dto3.Purchase{},
	"SetShoppingListNameRequest":            dto2.SetShoppingListName{},
	"UpdateShoppingListRequest":             dto2.SetShoppingList{},
	"ShoppingListVersionResponse":           dto1.SetShoppingList{},
	"ShoppingListSummary":                   dto1.ShoppingListInfo{},
	"ConfirmGoogleSubscriptionRequest":      dto8.ConfirmGoogleSubscription{},
	"Subscription":                          dto7.Subscription{},
	"Tag":                                   dto4.Tag{},
	"TagResponse":                           dto4.TagWithGroupName{},
	"TagsResponse":                          dto4.TagsAndGroups{},
}

type schema struct {
	Type       string             `yaml:"type"`
	Properties map[string]*schema `yaml:"properties"`
	Required   []string           `yaml:"required"`
	Nullable   bool               `yaml:"nullable"`
}

func jsonFields(typ reflect.Type) map[string]reflect.StructField {
	result := map[string]reflect.StructField{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			for k, v := range jsonFields(field.Type) {
				result[k] = v
			}
			continue
		}
		name := strings.Split(field.Tag.Get("json"), ",")[0]
		if name != "" && name != "-" {
			result[name] = field
		}
	}
	return result
}

func TestProviderDTOsMatchContract(t *testing.T) {
	var spec struct {
		Components struct {
			Schemas map[string]*schema `yaml:"schemas"`
		} `yaml:"components"`
	}
	if err := yaml.Unmarshal(contracts.OpenAPI, &spec); err != nil {
		t.Fatal(err)
	}
	// Manual transport DTOs are checked here; generated auth unions and response
	// envelopes are exercised by JSON conformance tests in the owning handlers.
	for name, value := range providerDTOs {
		t.Run(name, func(t *testing.T) {
			shape := spec.Components.Schemas[name]
			if shape == nil {
				t.Fatal("Schema missing from contract")
			}
			fields := jsonFields(reflect.TypeOf(value))
			if len(fields) != len(shape.Properties) {
				t.Errorf("Property count differs: Go %d, contract %d", len(fields), len(shape.Properties))
			}
			for jsonName, field := range fields {
				property := shape.Properties[jsonName]
				if property == nil {
					t.Errorf("Go JSON field %s missing from contract", jsonName)
					continue
				}
				if field.Type.Kind() == reflect.Pointer && !property.Nullable && !strings.Contains(field.Tag.Get("json"), "omitempty") {
					t.Errorf("Go pointer %s must allow null", jsonName)
				}
				if field.Tag.Get("binding") == "required" {
					found := false
					for _, required := range shape.Required {
						if required == jsonName {
							found = true
						}
					}
					if !found {
						t.Errorf("Required Go field %s is optional in contract", jsonName)
					}
				}
			}
		})
	}
}
