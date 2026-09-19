package auth

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"google.golang.org/grpc"
)

type stepClient struct {
	flowClient
	started *api.StartAuthenticationStepRequest
}

func (f *stepClient) StartAuthenticationStep(_ context.Context, in *api.StartAuthenticationStepRequest, _ ...grpc.CallOption) (*api.AuthenticationStep, error) {
	f.started = in
	f.calls++
	return &api.AuthenticationStep{Id: testID, Type: in.Type, Status: "pending", ExpirationTimestamp: state().ExpirationTimestamp}, nil
}

func TestStartStepPreservesTypeAndFlowBinding(t *testing.T) {
	f := &stepClient{}
	h := testHandler(&f.flowClient)
	h.service.Authentication = f
	w := call(h.StartAuthenticationStep, `{"type":"registration"}`, false)
	if w.Code != 201 || f.started.Type != "registration" || f.started.Flow.AuthenticationId != testID || f.started.Flow.FlowToken != "flow-secret" {
		t.Fatalf("lost typed flow: %d", w.Code)
	}
	if w.Header().Get("Location") != "/v1/authentications/"+testID+"/steps/"+testID {
		t.Fatal("incorrect step location")
	}
	validateResponse(t, w, "AuthenticationStep")
	for _, body := range []string{`{"type":"passkeyRegistration"}`, `{"type":"registration","email":"a@example.com"}`, `{"method":"email"}`, `{"type":"passwordSetup","password":"secret"}`, `{"type":"googleVerification"}`, `{"type":"googleVerification","credentialType":"idToken","redirectUri":"https://example.com"}`} {
		w = call(h.StartAuthenticationStep, body, false)
		if w.Code != 400 || f.calls != 1 {
			t.Fatalf("invalid step start reached RPC: %d", w.Code)
		}
	}
}

// Exercise every contract example through the real body validator and RPC mapper.
// This also catches new union variants that have no downstream mapping.
func TestStepCompletionExamplesReachTypedRPC(t *testing.T) {
	spec, err := authSpec()
	if err != nil {
		t.Fatal(err)
	}
	op := spec.Paths.Value("/v1/authentications/{id}/steps/{stepId}/completion").Post
	proofs := map[string]string{
		"registration": "registration", "passwordSetup": "password_setup",
		"passwordVerification": "password", "emailVerification": "email",
		"totpVerification": "totp", "backupCodeVerification": "backup_code",
		"googleVerification": "google", "vkVerification": "vk",
		"passkeyVerification": "passkey",
	}
	seen := map[string]bool{}
	for name, example := range op.RequestBody.Value.Content["application/json"].Examples {
		t.Run(name, func(t *testing.T) {
			body, err := json.Marshal(example.Value.Value)
			if err != nil {
				t.Fatal(err)
			}
			var value map[string]any
			if err := json.Unmarshal(body, &value); err != nil {
				t.Fatal(err)
			}
			kind := value["type"].(string)
			f := &flowClient{}
			w := call(testHandler(f).CompleteAuthenticationStep, string(body), false)
			if w.Code != 200 || f.attempt == nil {
				t.Fatalf("completion failed: %d", w.Code)
			}
			message := f.attempt.ProtoReflect()
			field := message.WhichOneof(message.Descriptor().Oneofs().ByName("proof"))
			if field == nil || string(field.Name()) != proofs[kind] {
				t.Fatalf("wrong RPC variant for %s: %v", kind, field)
			}
			if f.attempt.Step.StepId != testID || f.attempt.Step.Flow.FlowToken != "flow-secret" {
				t.Fatal("step lost flow binding")
			}
			seen[kind] = true
		})
	}
	if len(seen) != len(proofs) {
		t.Fatalf("covered %d of %d variants", len(seen), len(proofs))
	}
}

func TestStepCompletionRejectsLegacyAndCrossVariantFields(t *testing.T) {
	for _, body := range []string{
		`{"type":"passkeyRegistration","name":"Phone","credential":{}}`,
		`{"method":"password","password":"secret"}`,
		`{"type":"registration","registration":{"email":"a@example.com"}}`,
		`{"type":"registration","email":"a@example.com","username":"chef"}`,
		`{"type":"registration","email":"a@example.com","password":"secret-password"}`,
		`{"type":"passwordSetup","password":"secret-password","email":"a@example.com"}`,
		`{"type":"emailVerification","code":"123456","password":"secret"}`,
		`{"type":"emailVerification","code":123456}`,
		`{"type":"futureMethod","code":"123456"}`,
	} {
		f := &flowClient{}
		w := call(testHandler(f).CompleteAuthenticationStep, body, false)
		if w.Code != 400 || f.calls != 0 {
			t.Fatalf("invalid union reached RPC: status %d", w.Code)
		}
	}
}

func jsonRecorder(value any) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.JSON(200, value)
	return w
}

func TestReadSerializersNeverReplaySecrets(t *testing.T) {
	s := state()
	s.FlowToken = "flow-secret"
	s.ConfirmationToken = "grant-secret"
	s.ConfirmationExpirationTimestamp = s.ExpirationTimestamp
	s.Steps = []*api.AuthenticationStep{{Id: testID, Type: "googleVerification", Status: "pending", ExpirationTimestamp: s.ExpirationTimestamp, Url: "https://example.com/oauth?state=secret"}}
	w := jsonRecorder(stateJSON(s, false, false))
	validateResponse(t, w, "AuthenticationResponse")
	for _, value := range []string{"flowToken", "confirmationToken", "confirmationExpirationTimestamp", "url", "secret"} {
		if strings.Contains(w.Body.String(), value) {
			t.Fatalf("read leaked %s", value)
		}
	}
	for _, kind := range []string{"googleVerification", "vkVerification"} {
		s.Steps[0].Type = kind
		w = jsonRecorder(stepJSON(s.Steps[0], false))
		validateResponse(t, w, "AuthenticationStep")
		if strings.Contains(w.Body.String(), "url") {
			t.Fatal("step read leaked OAuth URL")
		}
		w = jsonRecorder(stepJSON(s.Steps[0], true))
		validateResponse(t, w, "AuthenticationStep")
		if !strings.Contains(w.Body.String(), "url") {
			t.Fatal("step creation omitted OAuth URL")
		}
	}
	s.Steps[0].Type, s.Steps[0].Url, s.Steps[0].Nonce = "googleVerification", "", "nonce-secret"
	for _, creation := range []bool{false, true} {
		w = jsonRecorder(stepJSON(s.Steps[0], creation))
		validateResponse(t, w, "AuthenticationStep")
		if strings.Contains(w.Body.String(), "nonce-secret") != creation {
			t.Fatal("wrong nonce disclosure context")
		}
	}
	w = jsonRecorder(stateJSON(s, false, true))
	if strings.Contains(w.Body.String(), "flowToken") || strings.Contains(w.Body.String(), "nonce-secret") || !strings.Contains(w.Body.String(), "confirmationToken") {
		t.Fatal("wrong completion disclosure context")
	}
}

func TestVerifiedExistingAccountConflictPreserved(t *testing.T) {
	f := &flowClient{completionError: fail.CreateGrpcConflict("account_exists", "account already exists")}
	w := call(testHandler(f).CompleteAuthenticationStep, `{"type":"emailVerification","code":"012345"}`, false)
	if w.Code != 409 || !strings.Contains(w.Body.String(), "account_exists") {
		t.Fatalf("lost verified account conflict: %d", w.Code)
	}
}
