package response

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"net/http"
)

func Unauthorized(c *gin.Context, err error) {
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
	c.AbortWithStatusJSON(response.Code, response)
}

func FailGrpc(c *gin.Context, err error) {
	response := fail.ParseGrpc(err)
	c.AbortWithStatusJSON(response.Code, response)
}
