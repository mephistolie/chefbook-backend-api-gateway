package auth

import (
	"context"
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-api-gateway/contracts"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/tokens/access"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type flowClient struct {
	api.AuthenticationServiceClient
	created         *api.CreateAuthenticationRequest
	attempt         *api.CompleteAuthenticationStepRequest
	session         *api.CreateAuthenticatedSessionRequest
	revoke          *api.RevokeAuthSessionRequest
	password        *api.SetAccountPasswordRequest
	completionError error
	calls           int
}

const testID = "c3d9f032-89d1-4d4c-a2d5-cffb102e87aa"

func state() *api.AuthenticationState {
	return &api.AuthenticationState{Id: testID, Purpose: "signIn", Status: "pending", ExpirationTimestamp: timestamppb.New(time.Now().Add(time.Minute)), Options: []string{"passwordVerification"}}
}
func (f *flowClient) CreateAuthentication(_ context.Context, in *api.CreateAuthenticationRequest, _ ...grpc.CallOption) (*api.AuthenticationState, error) {
	f.created = in
	f.calls++
	s := state()
	s.FlowToken = "opaque-flow"
	return s, nil
}
func (f *flowClient) CompleteAuthenticationStep(_ context.Context, in *api.CompleteAuthenticationStepRequest, _ ...grpc.CallOption) (*api.CompleteAuthenticationStepResponse, error) {
	f.attempt = in
	f.calls++
	s := state()
	if f.completionError != nil {
		return nil, f.completionError
	}
	return &api.CompleteAuthenticationStepResponse{Authentication: s, Step: &api.AuthenticationStep{Id: testID, Type: "passwordVerification", Status: "completed", ExpirationTimestamp: s.ExpirationTimestamp}}, nil
}
func (f *flowClient) CreateSession(_ context.Context, in *api.CreateAuthenticatedSessionRequest, _ ...grpc.CallOption) (*api.AuthTokens, error) {
	f.session = in
	f.calls++
	return &api.AuthTokens{UserId: testID, SessionId: 7, AccessToken: "access", RefreshToken: "refresh", ExpirationTimestamp: timestamppb.New(time.Now().Add(time.Hour))}, nil
}
func (f *flowClient) RevokeSessions(_ context.Context, in *api.RevokeAuthSessionRequest, _ ...grpc.CallOption) (*emptypb.Empty, error) {
	f.revoke = in
	f.calls++
	return &emptypb.Empty{}, nil
}
func (f *flowClient) SetPassword(_ context.Context, in *api.SetAccountPasswordRequest, _ ...grpc.CallOption) (*emptypb.Empty, error) {
	f.password = in
	f.calls++
	return &emptypb.Empty{}, nil
}
func testHandler(f *flowClient) *Handler {
	domain := "chefbook.test"
	return NewHandler(&service.Auth{Authentication: f}, config.Domains{Frontend: &domain})
}
func call(h gin.HandlerFunc, body string, authenticated bool) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/:id/:stepId", func(c *gin.Context) {
		if authenticated {
			request.PutUserPayload(c, access.Payload{UserId: uuid.MustParse(testID), SessionID: 7})
		}
		h(c)
	})
	req := httptest.NewRequest("POST", "/"+testID+"/"+testID, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Flow-Token", "flow-secret")
	req.Header.Set("Reauthentication-Token", "action-grant")
	req.Header.Set("Authorization", "Bearer refresh-secret")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func validateResponse(t *testing.T, w *httptest.ResponseRecorder, schema string) {
	t.Helper()
	spec, e := openapi3.NewLoader().LoadFromData(contracts.OpenAPI)
	if e != nil {
		t.Fatal(e)
	}
	var v any
	if e = json.Unmarshal(w.Body.Bytes(), &v); e != nil {
		t.Fatal(e)
	}
	if e = spec.Components.Schemas[schema].Value.VisitJSON(v); e != nil {
		t.Fatalf("response violates %s: %v", schema, e)
	}
}
func TestAuthenticationDTOAndSessionSeparation(t *testing.T) {
	f := &flowClient{}
	h := testHandler(f)
	w := call(h.CreateAuthentication, `{"purpose":{"type":"signUp"}}`, false)
	if w.Code != 201 || f.created.Purpose != "signUp" || f.created.Principal != nil {
		t.Fatalf("status=%d body=%s input=%v", w.Code, w.Body, f.created)
	}
	validateResponse(t, w, "AuthenticationResponse")
	w = call(h.CreateSession, `{"authenticationToken":"grant"}`, false)
	if w.Code != 201 || f.session.AuthenticationToken != "grant" || w.Header().Get("Location") != "/v1/sessions/7" {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	validateResponse(t, w, "TokensResponse")
}
func TestRejectAmbiguousAndLegacyBodiesBeforeRPC(t *testing.T) {
	for _, body := range []string{`{"purpose":{"type":"signUp"},"registration":{"email":"a@example.com","password":"x","provider":"google"}}`, `{"purpose":{"type":"signUp"},"registration":{"email":"a@example.com"}}`, `{"purpose":{"type":"signIn"},"login":"a@example.com"}`, `{"purpose":{"type":"signIn"},"client":{}}`} {
		f := &flowClient{}
		w := call(testHandler(f).CreateAuthentication, body, false)
		if w.Code != 400 || f.calls != 0 {
			t.Fatalf("accepted invalid request: %s status=%d", body, w.Code)
		}
	}
	f := &flowClient{}
	w := call(testHandler(f).CreateSession, `{"method":"password","login":"a@example.com","password":"x"}`, false)
	if w.Code != 400 || f.calls != 0 {
		t.Fatal("legacy credential session accepted")
	}
}
func TestProofCarriesOriginAndDoesNotUseAccountSelector(t *testing.T) {
	f := &flowClient{}
	h := testHandler(f)
	w := call(h.CompleteAuthenticationStep, `{"type":"passwordVerification","password":"secret"}`, true)
	if w.Code != 200 || f.attempt.GetPassword().Password != "secret" || f.attempt.Step.Flow.Principal.SessionId != 7 || f.attempt.Step.Flow.FlowToken != "flow-secret" {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	validateResponse(t, w, "CompleteAuthenticationStepResponse")
	w = call(h.CompleteAuthenticationStep, `{"type":"googleVerification","idToken":"x","code":"y","state":"z"}`, true)
	if w.Code != 400 || f.calls != 1 {
		t.Fatal("mixed OAuth proofs accepted")
	}
}
func TestActionGrantAndBodylessLogout(t *testing.T) {
	f := &flowClient{}
	h := testHandler(f)
	w := call(h.ChangePassword, `{"newPassword":"new-password"}`, true)
	if w.Code != 204 || w.Body.Len() != 0 || f.password.Action.ConfirmationToken != "action-grant" || f.password.Action.Principal.SessionId != 7 {
		t.Fatalf("status=%d body=%s", w.Code, w.Body)
	}
	w = call(func(c *gin.Context) { h.EndSession(c, 7) }, "", false)
	if w.Code != 204 || f.revoke.RefreshToken != "refresh-secret" || f.revoke.GetSessionId() != 7 {
		t.Fatal("logout lost path/refresh binding")
	}
	w = call(h.EndSessions, `{"ids":[7]}`, true)
	if w.Code != 400 {
		t.Fatal("legacy filter silently became delete all")
	}
}
