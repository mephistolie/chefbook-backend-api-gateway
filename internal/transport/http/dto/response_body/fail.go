package response_body

import (
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"net/http"
)

var (
	InvalidBody = fail.Response{Code: http.StatusBadRequest, ErrorType: fail.TypeInvalidBody, Message: "invalid body"}
)
