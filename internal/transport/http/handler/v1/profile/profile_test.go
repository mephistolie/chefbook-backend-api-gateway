package profile

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-common/tokens/access"
	api "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
	"google.golang.org/grpc"
)

type profileClient struct {
	api.ProfileServiceClient
	request *api.GetProfileRequest
}

func (c *profileClient) GetProfile(_ context.Context, req *api.GetProfileRequest, _ ...grpc.CallOption) (*api.GetProfileResponse, error) {
	c.request = req
	return &api.GetProfileResponse{Id: req.ProfileId}, nil
}

func TestProfileTargets(t *testing.T) {
	gin.SetMode(gin.TestMode)
	requester := uuid.New()
	other := uuid.New().String()
	for _, tc := range []struct{ name, path, id, username string }{
		{"self", "/v1/profile", requester.String(), ""},
		{"public ID", "/v1/profiles/" + other, other, ""},
		{"public username", "/v1/profiles/alice", "", "alice"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			client := &profileClient{}
			h := NewHandler(nil, nil, &service.Profile{ProfileServiceClient: client})
			r := gin.New()
			r.Use(func(c *gin.Context) { request.PutUserPayload(c, access.Payload{UserId: requester}) })
			r.GET("/v1/profile", h.GetProfile)
			r.GET("/v1/profiles/:profile_id", h.GetPublicProfile)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, httptest.NewRequest("GET", tc.path, nil))
			if w.Code != 200 || client.request == nil {
				t.Fatalf("status=%d body=%s", w.Code, w.Body)
			}
			if client.request.ProfileId != tc.id || client.request.ProfileUsername != tc.username || client.request.RequesterId != requester.String() {
				t.Fatalf("unexpected target: %+v", client.request)
			}
		})
	}
}
