package v1

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/contract"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/middleware/auth"
	"io"
	"net/http"
	"strings"
)

type Router struct {
	handler        v1.Handler
	authMiddleware interface {
		AuthorizeUser(*gin.Context)
		AuthorizeDeletedUser(*gin.Context)
	}
}

func NewRouter(handler v1.Handler, middleware *auth.Middleware) *Router {
	return &Router{handler: handler, authMiddleware: middleware}
}
func (r *Router) Init(api *gin.RouterGroup) {
	contract.RegisterHandlersWithOptions(api, r, contract.GinServerOptions{Middlewares: []contract.MiddlewareFunc{r.authorize}, ErrorHandler: func(c *gin.Context, err error, status int) { response.Fail(c, response.InvalidBody) }})
}
func authPath(path string) bool {
	for _, prefix := range []string{"/v1/authentications", "/v1/account/", "/v1/sessions", "/v1/oauth/", "/v1/usernames/"} {
		if path == prefix || strings.HasPrefix(path, prefix+"/") || strings.HasSuffix(prefix, "/") && strings.HasPrefix(path, prefix) {
			return true
		}
	}
	return false
}

// Runs before rate limiting and generated path/header binding to cover error responses.
func AuthCacheHeaders(c *gin.Context) {
	if authPath(c.Request.URL.Path) {
		c.Header("Cache-Control", "no-store")
		c.Header("Pragma", "no-cache")
	}
}
func allowsPendingDeletion(path string) bool {
	return path == "/v1/account/deletion" || path == "/v1/sessions" || strings.HasPrefix(path, "/v1/sessions/") || strings.HasPrefix(path, "/v1/authentications")
}
func (r *Router) authorize(c *gin.Context) {
	AuthCacheHeaders(c)
	path := c.FullPath()
	// OpenAPI cannot express bearer requirements depending on a JSON discriminator
	// or on the saved purpose of a flow. The owning service checks the latter again.
	if path == "/v1/authentications" && c.Request.Method == http.MethodPost {
		body, e := io.ReadAll(io.LimitReader(c.Request.Body, 64*1024+1))
		if e != nil || len(body) > 64*1024 {
			response.Fail(c, response.InvalidBody)
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewReader(body))
		var b struct {
			Purpose struct {
				Type string `json:"type"`
			} `json:"purpose"`
		}
		if json.Unmarshal(body, &b) != nil {
			response.Fail(c, response.InvalidBody)
			return
		}
		if b.Purpose.Type == "signIn" || b.Purpose.Type == "signUp" {
			return
		}
		r.authMiddleware.AuthorizeDeletedUser(c)
		return
	}
	if strings.HasPrefix(path, "/v1/authentications/") {
		if c.GetHeader("Authorization") != "" {
			r.authMiddleware.AuthorizeDeletedUser(c)
		}
		return // Flow-Token and original principal AND binding are verified by auth RPC.
	}
	if path == "/v1/sessions/:id" && c.Request.Method == http.MethodDelete {
		parts := strings.Fields(c.GetHeader("Authorization"))
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") && strings.Count(parts[1], ".") == 2 {
			r.authMiddleware.AuthorizeDeletedUser(c)
		}
		return // Opaque refresh logout authenticates only that path session in auth.
	}
	if _, protected := c.Get(contract.BearerAuthScopes); !protected {
		return
	}
	if allowsPendingDeletion(path) {
		r.authMiddleware.AuthorizeDeletedUser(c)
	} else {
		r.authMiddleware.AuthorizeUser(c)
	}
}
