package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
)

func invalidProviderResult(c *gin.Context) {
	response.Fail(c, fail.Response{Code: 500, ErrorType: "invalid_provider_response", Message: "Invalid authentication service response"})
}
