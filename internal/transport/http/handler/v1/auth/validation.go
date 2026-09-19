package auth

import (
	"encoding/json"
	"io"
	"strings"
	"sync"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/contracts"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

var authSpec = sync.OnceValues(func() (*openapi3.T, error) { return openapi3.NewLoader().LoadFromData(contracts.OpenAPI) })

// Validate unions and required fields against the pinned source before decoding Go
// DTOs: code generators alone do not enforce oneOf/not or required boolean fields.
func bind(c *gin.Context, schema string, target any) bool {
	data, err := io.ReadAll(io.LimitReader(c.Request.Body, 64*1024+1))
	if err != nil || len(data) > 64*1024 {
		invalid(c)
		return false
	}
	var value any
	if json.Unmarshal(data, &value) != nil {
		invalid(c)
		return false
	}
	spec, err := authSpec()
	if err != nil {
		response.Fail(c, fail.Response{Code: 500, ErrorType: "unknown", Message: "contract unavailable"})
		return false
	}
	shape := spec.Components.Schemas[schema]
	if shape == nil || shape.Value.VisitJSON(value) != nil || json.Unmarshal(data, target) != nil {
		invalid(c)
		return false
	}
	return true
}
func invalid(c *gin.Context) {
	response.Fail(c, fail.Response{Code: 400, ErrorType: "invalid_request", Message: "invalid request body"})
}
func loginParts(login string) (email, username string) {
	login = strings.TrimSpace(login)
	if strings.Contains(login, "@") {
		return login, ""
	}
	return "", login
}
func noBody(c *gin.Context) bool {
	if c.Request.Body != nil {
		b, e := io.ReadAll(io.LimitReader(c.Request.Body, 1))
		if e != nil || len(b) != 0 {
			invalid(c)
			return false
		}
	}
	return true
}
