package response

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGrpcHTTPStatuses(t *testing.T) {
	gin.SetMode(gin.TestMode)
	for _, tc := range []struct {
		name   string
		err    error
		code   int
		reason string
	}{
		{"legacy missing", fail.CreateGrpcClient(fail.TypeNotFound, "missing"), 404, "not_found"},
		{"missing", fail.CreateGrpcNotFound(fail.TypeNotFound, "missing"), 404, "not_found"},
		{"version conflict", fail.CreateGrpcConflict("outdated_version", "stale"), 409, "outdated_version"},
		{"email occupied", fail.CreateGrpcClient("profile_exists", "exists"), 409, "profile_exists"},
		{"username occupied", fail.CreateGrpcClient("username_occupied", "exists"), 409, "username_occupied"},
		{"access denied", fail.CreateGrpcClient(fail.TypeAccessDenied, "denied"), 403, "access_denied"},
		{"invalid input", fail.GrpcInvalidBody, 400, "invalid_body"},
		{"credentials unchanged", fail.CreateGrpcClient("invalid_credentials", "invalid"), 400, "invalid_credentials"},
		{"unavailable", status.Error(codes.Unavailable, "offline"), 503, "unavailable"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			FailGrpc(c, tc.err)
			var body fail.Response
			if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			if w.Code != tc.code || body.ErrorType != tc.reason {
				t.Fatalf("status=%d body=%s", w.Code, w.Body)
			}
		})
	}
}
