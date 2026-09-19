package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/contracts"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/contract"
	"go.yaml.in/yaml/v3"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

type denyAuth struct{ deleted bool }

func (a *denyAuth) AuthorizeUser(c *gin.Context)        { a.deleted = false; c.AbortWithStatus(401) }
func (a *denyAuth) AuthorizeDeletedUser(c *gin.Context) { a.deleted = true; c.AbortWithStatus(401) }

type unreachableServer struct{ contract.ServerInterface }

func TestStepRoutesReplaceChallengeRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	contract.RegisterHandlers(engine, &unreachableServer{})
	routes := map[string]bool{}
	for _, route := range engine.Routes() {
		routes[route.Method+" "+route.Path] = true
		if strings.Contains(route.Path, "/challenges") || strings.Contains(route.Path, "/attempts") {
			t.Fatalf("legacy authentication route retained: %s", route.Path)
		}
	}
	for _, path := range []string{
		"POST /v1/authentications/:id/steps",
		"GET /v1/authentications/:id/steps/:stepId",
		"POST /v1/authentications/:id/steps/:stepId/completion",
	} {
		if !routes[path] {
			t.Fatalf("missing step route: %s", path)
		}
	}
}

func TestContractAuthAndRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var spec struct {
		Paths map[string]map[string]struct {
			OperationID string                `yaml:"operationId"`
			Security    []map[string][]string `yaml:"security"`
		} `yaml:"paths"`
	}
	if e := yaml.Unmarshal(contracts.OpenAPI, &spec); e != nil {
		t.Fatal(e)
	}
	auth := &denyAuth{}
	r := &Router{authMiddleware: auth}
	engine := gin.New()
	contract.RegisterHandlersWithOptions(engine, &unreachableServer{}, contract.GinServerOptions{Middlewares: []contract.MiddlewareFunc{r.authorize}})
	parameter := regexp.MustCompile(`\{[^}]+\}`)
	count := 0
	protected := 0
	for path, methods := range spec.Paths {
		for method, op := range methods {
			count++
			// Conditional public/flow/logout paths are tested separately, backend verifies saved context.
			if strings.HasPrefix(path, "/v1/authentications") || op.OperationID == "revokeSession" {
				continue
			}
			needsBearer := false
			for _, entry := range op.Security {
				if _, ok := entry["bearerAuth"]; ok {
					needsBearer = true
				}
			}
			if !needsBearer {
				continue
			}
			protected++
			concrete := parameter.ReplaceAllString(path, "c3d9f032-89d1-4d4c-a2d5-cffb102e87aa")
			w := httptest.NewRecorder()
			req := httptest.NewRequest(strings.ToUpper(method), concrete, nil)
			req.Header.Set("Reauthentication-Token", "one-use-grant")
			engine.ServeHTTP(w, req)
			if w.Code != 401 {
				t.Errorf("%s %s: %d", method, path, w.Code)
			}
			if auth.deleted != allowsPendingDeletion(path) {
				t.Errorf("deletion policy %s", path)
			}
		}
	}
	if len(engine.Routes()) != count || protected < 30 {
		t.Fatalf("coverage routes=%d expected=%d protected=%d", len(engine.Routes()), count, protected)
	}
}
func TestConditionalAuthenticationSecurity(t *testing.T) {
	for _, purpose := range []string{"signIn", "signUp", "passwordChange", "identityLink", "unknown"} {
		a := &denyAuth{}
		r := &Router{authMiddleware: a}
		engine := gin.New()
		engine.POST("/v1/authentications", r.authorize, func(c *gin.Context) { c.Status(201) })
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, httptest.NewRequest("POST", "/v1/authentications", strings.NewReader(`{"purpose":{"type":"`+purpose+`"}}`)))
		want := 401
		if purpose == "signIn" || purpose == "signUp" {
			want = 201
		}
		if w.Code != want {
			t.Errorf("%s: %d", purpose, w.Code)
		}
	}
}
func TestTokenResponsesNeverCacheEvenOnEarlyFailure(t *testing.T) {
	for _, path := range []string{"/v1/sessions", "/v1/sessions/1/tokens", "/v1/oauth/google", "/v1/oauth/vk", "/v1/account/totp", "/v1/authentications", "/v1/authentications/anything"} {
		engine := gin.New()
		engine.Use(AuthCacheHeaders, func(c *gin.Context) { c.AbortWithStatus(429) })
		engine.POST(path, func(c *gin.Context) { t.Fatal("reached handler") })
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, httptest.NewRequest("POST", path, nil))
		if w.Code != 429 || w.Header().Get("Cache-Control") != "no-store" || w.Header().Get("Pragma") != "no-cache" {
			t.Errorf("headers missing for %s", path)
		}
	}
}
