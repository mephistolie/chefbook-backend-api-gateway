package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/tokens/access"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http/httptest"
	"testing"
	"time"
)

type stateClient struct {
	api.AuthServiceClient
	info *api.GetAuthInfoResponse
	err  error
}

func (s stateClient) GetAuthInfo(context.Context, *api.GetAuthInfoRequest, ...grpc.CallOption) (*api.GetAuthInfoResponse, error) {
	return s.info, s.err
}
func TestLiveAccountStateOverridesOldJWT(t *testing.T) {
	gin.SetMode(gin.TestMode)
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name                                           string
		claimedDeleted, deleted, blocked, allowDeleted bool
		method, path                                   string
		upstream                                       error
		want                                           int
	}{
		{name: "deletion after issuance", deleted: true, method: "GET", path: "/v1/profile", want: 409},
		{name: "cancelled deletion", claimedDeleted: true, method: "GET", path: "/v1/profile", want: 204},
		{name: "blocked after issuance", blocked: true, method: "GET", path: "/v1/profile", want: 403},
		{name: "deleting can list sessions", deleted: true, allowDeleted: true, method: "GET", path: "/v1/sessions", want: 204},
		{name: "blocked cannot list sessions", blocked: true, allowDeleted: true, method: "GET", path: "/v1/sessions", want: 403},
		{name: "blocked can revoke sessions", blocked: true, allowDeleted: true, method: "DELETE", path: "/v1/sessions", want: 204},
		{name: "permanently deleted", method: "GET", path: "/v1/profile", upstream: status.Error(codes.NotFound, "missing"), want: 401},
		{name: "dependency offline", method: "GET", path: "/v1/profile", upstream: status.Error(codes.Unavailable, "offline"), want: 503},
	} {
		t.Run(tc.name, func(t *testing.T) {
			info := &api.GetAuthInfoResponse{IsActivated: true, IsBlocked: tc.blocked}
			if tc.deleted {
				info.DeletionTimestamp = timestamppb.New(time.Now().Add(time.Hour))
			}
			m := &Middleware{authService: &service.Auth{AuthServiceClient: stateClient{info: info, err: tc.upstream}, Authentication: liveSessionClient{}}, tokenParser: access.NewParserByKey(&key.PublicKey), keyUpdateTimestamp: time.Now(), keyUpdateInterval: time.Hour}
			token, err := access.NewProducerByKey(key).Produce(access.Payload{UserId: uuid.New(), SessionID: 7, Deleted: tc.claimedDeleted}, time.Hour)
			if err != nil {
				t.Fatal(err)
			}
			engine := gin.New()
			engine.Handle(tc.method, tc.path, func(c *gin.Context) { m.authorizeUser(c, tc.allowDeleted) }, func(c *gin.Context) { c.Status(204) })
			req := httptest.NewRequest(tc.method, tc.path, nil)
			req.Header.Set("Authorization", "Bearer "+token)
			w := httptest.NewRecorder()
			engine.ServeHTTP(w, req)
			if w.Code != tc.want {
				t.Fatalf("status=%d body=%s", w.Code, w.Body)
			}
			if w.Code == 401 && w.Header().Get("WWW-Authenticate") == "" {
				t.Fatal("missing challenge")
			}
		})
	}
}

type liveSessionClient struct {
	api.AuthenticationServiceClient
	err error
}

func (s liveSessionClient) ValidateSession(context.Context, *api.AuthPrincipal, ...grpc.CallOption) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.err
}
func TestRevokedSessionCannotUseStillValidJWT(t *testing.T) {
	gin.SetMode(gin.TestMode)
	key, e := rsa.GenerateKey(rand.Reader, 2048)
	if e != nil {
		t.Fatal(e)
	}
	m := &Middleware{authService: &service.Auth{Authentication: liveSessionClient{err: status.Error(codes.Unauthenticated, "revoked")}}, tokenParser: access.NewParserByKey(&key.PublicKey), keyUpdateTimestamp: time.Now(), keyUpdateInterval: time.Hour}
	token, e := access.NewProducerByKey(key).Produce(access.Payload{UserId: uuid.New(), SessionID: 7}, time.Hour)
	if e != nil {
		t.Fatal(e)
	}
	engine := gin.New()
	engine.GET("/profile", m.AuthorizeUser, func(c *gin.Context) { t.Fatal("revoked session reached handler") })
	req := httptest.NewRequest("GET", "/profile", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	if w.Code != 401 || w.Header().Get("WWW-Authenticate") == "" {
		t.Fatalf("status=%d headers=%v", w.Code, w.Header())
	}
}
