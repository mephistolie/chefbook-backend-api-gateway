package v1_test

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-api-gateway/contracts"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/encryption"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe"
	shopping "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/subscription"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-common/tokens/access"
	encapi "github.com/mephistolie/chefbook-backend-encryption/api/proto/implementation/v1"
	recipeapi "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	shoppingapi "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
	subapi "github.com/mephistolie/chefbook-backend-subscription/api/proto/implementation/v1"
	"google.golang.org/grpc"
)

func (recipeClient) GenerateRecipePicturesUploadLinks(context.Context, *recipeapi.GenerateRecipePicturesUploadLinksRequest, ...grpc.CallOption) (*recipeapi.GenerateRecipePicturesUploadLinksResponse, error) {
	return &recipeapi.GenerateRecipePicturesUploadLinksResponse{}, nil
}
func (shoppingClient) GetShoppingLists(context.Context, *shoppingapi.GetShoppingListsRequest, ...grpc.CallOption) (*shoppingapi.GetShoppingListsResponse, error) {
	return &shoppingapi.GetShoppingListsResponse{}, nil
}
func (shoppingClient) GetShoppingListUsers(context.Context, *shoppingapi.GetShoppingListUsersRequest, ...grpc.CallOption) (*shoppingapi.GetShoppingListUsersResponse, error) {
	return &shoppingapi.GetShoppingListUsersResponse{}, nil
}

type encryptionClient struct{ encapi.EncryptionServiceClient }

func (encryptionClient) GetRecipeKeyRequests(context.Context, *encapi.GetRecipeKeyRequestsRequest, ...grpc.CallOption) (*encapi.GetRecipeKeyRequestsResponse, error) {
	return &encapi.GetRecipeKeyRequestsResponse{}, nil
}

type subscriptionClient struct {
	subapi.SubscriptionServiceClient
}

func (subscriptionClient) GetProfileSubscriptions(context.Context, *subapi.GetProfileSubscriptionsRequest, ...grpc.CallOption) (*subapi.GetProfileSubscriptionsResponse, error) {
	return &subapi.GetProfileSubscriptionsResponse{}, nil
}
func TestEmptyCollectionEnvelopes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	domain := "chefbook.test"
	rh := recipe.NewHandler(&service.Recipe{RecipeServiceClient: recipeClient{}})
	sh := shopping.NewHandler(&service.ShoppingList{ShoppingListServiceClient: shoppingClient{}}, config.Domains{Frontend: &domain})
	eh := encryption.NewHandler(&service.Encryption{EncryptionServiceClient: encryptionClient{}})
	suh := subscription.NewHandler(&service.Subscription{SubscriptionServiceClient: subscriptionClient{}})
	spec, err := openapi3.NewLoader().LoadFromData(contracts.OpenAPI)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		schema, field, method, body string
		handler                     gin.HandlerFunc
	}{
		{"RecipePictureUploadLinksResponse", "uploads", "POST", `{"picturesCount":1}`, rh.GenerateRecipePicturesUploadLinks},
		{"RecipeKeyAccessRequestsResponse", "requests", "GET", "", eh.GetRecipeKeyRequests},
		{"ShoppingListsResponse", "shoppingLists", "GET", "", sh.GetShoppingLists},
		{"ShoppingListUsersResponse", "users", "GET", "", sh.GetShoppingListUsers},
		{"SubscriptionsResponse", "subscriptions", "GET", "", suh.GetSubscriptions},
	} {
		t.Run(tc.schema, func(t *testing.T) {
			e := gin.New()
			e.Use(func(c *gin.Context) { request.PutUserPayload(c, access.Payload{UserId: uuid.New()}) })
			e.Handle(tc.method, "/", tc.handler)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, "/", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			e.ServeHTTP(w, req)
			if w.Code != 200 {
				t.Fatalf("%d %s", w.Code, w.Body)
			}
			var body map[string]any
			if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			list, ok := body[tc.field].([]any)
			if !ok || len(list) != 0 {
				t.Fatalf("empty collection must be []: %s", w.Body)
			}
			if err := spec.Components.Schemas[tc.schema].Value.VisitJSON(body); err != nil {
				t.Fatal(err)
			}
		})
	}
}
