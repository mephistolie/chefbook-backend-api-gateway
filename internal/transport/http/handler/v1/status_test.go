package v1_test

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-api-gateway/contracts"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	auth "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth"
	recipe "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe"
	shopping "github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/shopping_list"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	authapi "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/mephistolie/chefbook-backend-common/tokens/access"
	recipeapi "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	shoppingapi "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
	"go.yaml.in/yaml/v3"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type recipeClient struct {
	recipeapi.RecipeServiceClient
	err error
}

func (c recipeClient) CreateRecipe(context.Context, *recipeapi.RecipeInput, ...grpc.CallOption) (*recipeapi.CreateRecipeResponse, error) {
	return &recipeapi.CreateRecipeResponse{RecipeId: "created-id", Version: 1}, c.err
}
func (c recipeClient) CreateCollection(context.Context, *recipeapi.CreateCollectionRequest, ...grpc.CallOption) (*recipeapi.CreateCollectionResponse, error) {
	return &recipeapi.CreateCollectionResponse{CollectionId: "created-id"}, c.err
}

type shoppingClient struct {
	shoppingapi.ShoppingListServiceClient
	err error
}

func (c shoppingClient) CreateSharedShoppingList(context.Context, *shoppingapi.CreateSharedShoppingListRequest, ...grpc.CallOption) (*shoppingapi.CreateSharedShoppingListResponse, error) {
	return &shoppingapi.CreateSharedShoppingListResponse{ShoppingListId: "created-id"}, c.err
}

type authClient struct {
	authapi.AuthenticationServiceClient
	err error
}

func (c authClient) RequestDeletion(context.Context, *authapi.RequestAccountDeletionRequest, ...grpc.CallOption) (*authapi.AccountDeletion, error) {
	return &authapi.AccountDeletion{DeletionTimestamp: timestamppb.New(time.Now().Add(time.Hour))}, c.err
}

func TestMutationResponseContract(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var spec struct {
		Paths map[string]map[string]struct {
			Responses map[string]any `yaml:"responses"`
		} `yaml:"paths"`
	}
	if err := yaml.Unmarshal(contracts.OpenAPI, &spec); err != nil {
		t.Fatal(err)
	}
	for _, upstreamErr := range []error{nil, fail.CreateGrpcConflict("outdated_version", "conflict")} {
		domain := "chefbook.test"
		rh := recipe.NewHandler(&service.Recipe{RecipeServiceClient: recipeClient{err: upstreamErr}})
		sh := shopping.NewHandler(&service.ShoppingList{ShoppingListServiceClient: shoppingClient{err: upstreamErr}}, config.Domains{Frontend: &domain})
		ph := auth.NewHandler(&service.Auth{Authentication: authClient{err: upstreamErr}}, config.Domains{Frontend: &domain})
		for _, tc := range []struct {
			path, body, location, field string
			code                        int
			handler                     gin.HandlerFunc
		}{
			{"/v1/recipes", `{"name":"Recipe","ingredients":[],"cooking":[]}`, "/v1/recipes/created-id", "id", 201, rh.CreateRecipe},
			{"/v1/collections", `{"name":"Collection"}`, "/v1/collections/created-id", "id", 201, rh.AddCollection},
			{"/v1/shopping-lists", `{}`, "/v1/shopping-lists/created-id", "id", 201, sh.CreateSharedShoppingList},
			{"/v1/account/deletion", `{"deleteSharedData":false}`, "/v1/account/deletion", "deletionTimestamp", 202, ph.DeleteProfile},
		} {
			t.Run(tc.path+strconv.FormatBool(upstreamErr != nil), func(t *testing.T) {
				r := gin.New()
				r.Use(func(c *gin.Context) { request.PutUserPayload(c, access.Payload{UserId: uuid.New()}) })
				r.POST(tc.path, tc.handler)
				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", tc.path, strings.NewReader(tc.body))
				req.Header.Set("Content-Type", "application/json")
				r.ServeHTTP(w, req)
				want := tc.code
				if upstreamErr != nil {
					want = 409
				}
				if w.Code != want {
					t.Fatalf("status=%d body=%s", w.Code, w.Body)
				}
				if upstreamErr != nil {
					if w.Header().Get("Location") != "" {
						t.Fatal("error advertised successful creation")
					}
					return
				}
				if _, ok := spec.Paths[tc.path]["post"].Responses[strconv.Itoa(w.Code)]; !ok {
					t.Fatal("status missing from OpenAPI")
				}
				if w.Header().Get("Location") != tc.location {
					t.Fatalf("Location=%s", w.Header().Get("Location"))
				}
				var body map[string]any
				if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
					t.Fatal(err)
				}
				if _, ok := body[tc.field]; !ok {
					t.Fatalf("missing response field %s: %s", tc.field, w.Body)
				}
			})
		}
	}
}
