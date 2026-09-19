package response

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

const (
	TypeProfileDeleting = "account_deleting"
)

var (
	InvalidBody = fail.Response{Code: http.StatusBadRequest, ErrorType: fail.TypeInvalidBody, Message: "invalid body"}

	ProfileDeleting = fail.Response{Code: http.StatusConflict, ErrorType: TypeProfileDeleting, Message: "profile is being deleted"}
)

func Unauthorized(c *gin.Context, err error) {
	c.Header("WWW-Authenticate", "Bearer")
	response := fail.Response{
		Code:      http.StatusUnauthorized,
		ErrorType: fail.TypeUnauthorized,
		Message:   err.Error(),
	}
	c.AbortWithStatusJSON(response.Code, response)
}

func Unknown(c *gin.Context, err error) {
	response := fail.Response{
		Code:      http.StatusInternalServerError,
		ErrorType: fail.TypeUnknown,
		Message:   err.Error(),
	}
	c.AbortWithStatusJSON(response.Code, response)
}

func Fail(c *gin.Context, response fail.Response) {
	if response.Code == http.StatusTooManyRequests && c.Writer.Header().Get("Retry-After") == "" {
		c.Header("Retry-After", "60")
	}
	if response.Code == http.StatusUnauthorized {
		c.Header("WWW-Authenticate", `Bearer error="invalid_token"`)
	}
	response = normalizeFailure(response)
	c.AbortWithStatusJSON(response.Code, response)
}

func FailGrpc(c *gin.Context, err error) {
	response := fail.ParseGrpc(err)
	switch status.Code(err) {
	case codes.DeadlineExceeded:
		response.Code = 503
		response.ErrorType = "unavailable"
		response.Message = "service unavailable"
	case codes.ResourceExhausted:
		response.Code = 429
		response.ErrorType = "rate_limited"
		response.Message = "too many requests"
	}
	Fail(c, response)
}

// Older services encode some semantic errors as InvalidArgument. Normalize only
// those known reasons at the HTTP boundary, preserving their public JSON bodies.
func normalizeFailure(response fail.Response) fail.Response {
	if response.Code == http.StatusBadRequest {
		switch response.ErrorType {
		case fail.TypeNotFound:
			response.Code = http.StatusNotFound
		case fail.TypeConflict, "profile_exists", "username_occupied", "account_occupied", "few_sign_in_methods", TypeProfileDeleting:
			response.Code = http.StatusConflict
		case fail.TypeAccessDenied, fail.TypePremiumRequired:
			response.Code = http.StatusForbidden
		}
	}
	return response
}
