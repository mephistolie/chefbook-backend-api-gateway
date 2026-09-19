package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/tokens/access"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"strings"
)

// Optional identity comes only from verified middleware, never request JSON.
func principal(c *gin.Context) *api.AuthPrincipal {
	raw, exists := c.Get("userPayload")
	if !exists {
		return nil
	}
	p, ok := raw.(access.Payload)
	if !ok {
		return nil
	}
	return &api.AuthPrincipal{AccountId: p.UserId.String(), SessionId: p.SessionID}
}
func flow(c *gin.Context) *api.FlowContext {
	return &api.FlowContext{AuthenticationId: c.Param("id"), FlowToken: c.GetHeader("Flow-Token"), Principal: principal(c)}
}
func stepContext(c *gin.Context) *api.AuthenticationStepContext {
	return &api.AuthenticationStepContext{Flow: flow(c), StepId: c.Param("stepId")}
}
func sensitive(c *gin.Context) *api.SensitiveAction {
	return &api.SensitiveAction{Principal: principal(c), ConfirmationToken: c.GetHeader("Reauthentication-Token")}
}
func failed(c *gin.Context, err error) bool {
	if err != nil {
		response.FailGrpc(c, err)
		return true
	}
	return false
}
func timestamp(t *timestamppb.Timestamp) any {
	if t == nil {
		return nil
	}
	return t.AsTime()
}
func emptyStrings(values []string) []string {
	if values == nil {
		return []string{}
	}
	return values
}
func stateJSON(s *api.AuthenticationState, creation, mutation bool) gin.H {
	options := []gin.H{}
	for _, kind := range s.Options {
		options = append(options, gin.H{"type": kind})
	}
	steps := []gin.H{}
	for _, step := range s.Steps {
		steps = append(steps, stepJSON(step, false))
	}
	out := gin.H{"id": s.Id, "purpose": gin.H{"type": s.Purpose}, "status": s.Status, "expirationTimestamp": timestamp(s.ExpirationTimestamp), "next": gin.H{"options": options}, "steps": steps}
	if creation && s.FlowToken != "" {
		out["flowToken"] = s.FlowToken
	}
	if mutation && s.ConfirmationToken != "" {
		out["confirmationToken"] = s.ConfirmationToken
		out["confirmationExpirationTimestamp"] = timestamp(s.ConfirmationExpirationTimestamp)
	}
	return out
}
func stepJSON(s *api.AuthenticationStep, creation bool) gin.H {
	out := gin.H{"id": s.Id, "type": s.Type, "status": s.Status, "expirationTimestamp": timestamp(s.ExpirationTimestamp)}
	switch s.Type {
	case "emailVerification":
		out["codeLength"] = s.CodeLength
	case "googleVerification", "vkVerification":
		if creation && s.Url != "" {
			out["url"] = s.Url
		}
		if creation && s.Type == "googleVerification" && s.Nonce != "" {
			out["nonce"] = s.Nonce
		}
	case "passkeyVerification":
		if len(s.PublicKeyOptions) > 0 {
			out["publicKey"] = json.RawMessage(s.PublicKeyOptions)
		}
	}
	return out
}
func tokensJSON(s *api.AuthTokens) gin.H {
	restrictions := []gin.H{}
	for _, r := range s.Restrictions {
		restrictions = append(restrictions, gin.H{"type": r.Type, "deletionTimestamp": timestamp(r.DeletionTimestamp), "deleteSharedData": r.DeleteSharedData})
	}
	return gin.H{"userId": s.UserId, "sessionId": s.SessionId, "accessToken": s.AccessToken, "refreshToken": s.RefreshToken, "expirationTimestamp": timestamp(s.ExpirationTimestamp), "restrictions": restrictions}
}
func (h *Handler) CreateAuthentication(c *gin.Context) {
	var b struct {
		Purpose struct {
			Type string `json:"type"`
		} `json:"purpose"`
	}
	if !bind(c, "StartAuthenticationRequest", &b) {
		return
	}
	in := &api.CreateAuthenticationRequest{Purpose: b.Purpose.Type, Principal: principal(c)}
	s, e := h.service.Authentication.CreateAuthentication(c, in)
	if failed(c, e) {
		return
	}
	c.Header("Location", "/v1/authentications/"+s.Id)
	c.JSON(201, stateJSON(s, true, true))
}
func (h *Handler) GetAuthentication(c *gin.Context) {
	s, e := h.service.Authentication.GetAuthentication(c, flow(c))
	if failed(c, e) {
		return
	}
	c.JSON(200, stateJSON(s, false, false))
}
func (h *Handler) StartAuthenticationStep(c *gin.Context) {
	var b struct {
		Type           string `json:"type"`
		RedirectUri    string `json:"redirectUri"`
		CredentialType string `json:"credentialType"`
	}
	if !bind(c, "StartAuthenticationStepRequest", &b) {
		return
	}
	s, e := h.service.Authentication.StartAuthenticationStep(c, &api.StartAuthenticationStepRequest{Flow: flow(c), Type: b.Type, RedirectUri: b.RedirectUri, CredentialType: b.CredentialType})
	if failed(c, e) {
		return
	}
	c.Header("Location", fmt.Sprintf("/v1/authentications/%s/steps/%s", c.Param("id"), s.Id))
	c.JSON(201, stepJSON(s, true))
}
func (h *Handler) GetAuthenticationStep(c *gin.Context) {
	s, e := h.service.Authentication.GetAuthenticationStep(c, stepContext(c))
	if failed(c, e) {
		return
	}
	c.JSON(200, stepJSON(s, false))
}
func (h *Handler) CompleteAuthenticationStep(c *gin.Context) {
	var b struct {
		Type       string          `json:"type"`
		Email      string          `json:"email"`
		Password   string          `json:"password"`
		Login      string          `json:"login"`
		Code       string          `json:"code"`
		State      string          `json:"state"`
		IdToken    string          `json:"idToken"`
		Credential json.RawMessage `json:"credential"`
	}
	if !bind(c, "CompleteAuthenticationStepRequest", &b) {
		return
	}
	in := &api.CompleteAuthenticationStepRequest{Step: stepContext(c)}
	switch b.Type {
	case "registration":
		in.Proof = &api.CompleteAuthenticationStepRequest_Registration{Registration: &api.RegistrationStepData{Email: b.Email}}
	case "passwordSetup":
		in.Proof = &api.CompleteAuthenticationStepRequest_PasswordSetup{PasswordSetup: &api.PasswordSetupStepData{Password: b.Password}}
	case "passwordVerification":
		in.Proof = &api.CompleteAuthenticationStepRequest_Password{Password: &api.PasswordChallengeProof{Password: b.Password, Login: b.Login}}
	case "emailVerification":
		in.Proof = &api.CompleteAuthenticationStepRequest_Email{Email: &api.CodeChallengeProof{Code: b.Code}}
	case "totpVerification":
		in.Proof = &api.CompleteAuthenticationStepRequest_Totp{Totp: &api.CodeChallengeProof{Code: b.Code}}
	case "backupCodeVerification":
		in.Proof = &api.CompleteAuthenticationStepRequest_BackupCode{BackupCode: &api.CodeChallengeProof{Code: b.Code}}
	case "googleVerification":
		in.Proof = &api.CompleteAuthenticationStepRequest_Google{Google: &api.OAuthChallengeProof{Code: b.Code, State: b.State, IdToken: b.IdToken}}
	case "vkVerification":
		in.Proof = &api.CompleteAuthenticationStepRequest_Vk{Vk: &api.OAuthChallengeProof{Code: b.Code, State: b.State}}
	case "passkeyVerification":
		in.Proof = &api.CompleteAuthenticationStepRequest_Passkey{Passkey: &api.PasskeyChallengeProof{Credential: b.Credential}}
	default:
		invalid(c)
		return
	}
	s, e := h.service.Authentication.CompleteAuthenticationStep(c, in)
	if failed(c, e) {
		return
	}
	c.JSON(200, gin.H{"step": stepJSON(s.Step, false), "authentication": stateJSON(s.Authentication, false, true)})
}
func (h *Handler) CreateSession(c *gin.Context) {
	var b struct {
		Token string `json:"authenticationToken"`
	}
	if !bind(c, "CreateSessionRequest", &b) {
		return
	}
	s, e := h.service.Authentication.CreateSession(c, &api.CreateAuthenticatedSessionRequest{AuthenticationToken: b.Token, Ip: c.ClientIP(), UserAgent: c.Request.UserAgent()})
	if failed(c, e) {
		return
	}
	c.Header("Location", fmt.Sprintf("/v1/sessions/%d", s.SessionId))
	c.JSON(201, tokensJSON(s))
}
func (h *Handler) RefreshSession(c *gin.Context) {
	var b struct {
		Token string `json:"refreshToken"`
	}
	if !bind(c, "RefreshSessionTokensRequest", &b) {
		return
	}
	id, e := strconv.ParseInt(c.Param("id"), 10, 64)
	if e != nil || id <= 0 {
		invalid(c)
		return
	}
	s, e := h.service.Authentication.RefreshSession(c, &api.RotateSessionRequest{SessionId: id, RefreshToken: b.Token, Ip: c.ClientIP(), UserAgent: c.Request.UserAgent()})
	if failed(c, e) {
		return
	}
	c.JSON(200, tokensJSON(s))
}
func (h *Handler) GetSessions(c *gin.Context) {
	s, e := h.service.Authentication.GetSessions(c, principal(c))
	if failed(c, e) {
		return
	}
	sessions := []gin.H{}
	for _, v := range s.Sessions {
		sessions = append(sessions, gin.H{"id": v.Id, "ip": v.Ip, "lastRefreshTimestamp": timestamp(v.LastRefreshTimestamp), "location": nil, "client": sessionClient(v.UserAgent)})
	}
	c.JSON(200, gin.H{"sessions": sessions})
}
func sessionClient(ua string) gin.H {
	platform, kind := "unknown", "unknown"
	u := strings.ToLower(ua)
	switch {
	case strings.Contains(u, "android"):
		platform = "android"
	case strings.Contains(u, "iphone") || strings.Contains(u, "ipad"):
		platform = "ios"
	case strings.Contains(u, "windows"):
		platform = "windows"
	case strings.Contains(u, "macintosh") || strings.Contains(u, "macos"):
		platform = "macos"
	case strings.Contains(u, "linux"):
		platform = "linux"
	}
	switch {
	case strings.Contains(u, "chefbook"):
		kind = "app"
	case strings.Contains(u, "mozilla/"):
		kind = "browser"
	}
	return gin.H{"name": nil, "platform": platform, "type": kind}
}
func (h *Handler) EndSessions(c *gin.Context) {
	if !noBody(c) {
		return
	}
	_, e := h.service.Authentication.RevokeSessions(c, &api.RevokeAuthSessionRequest{Principal: principal(c)})
	if !failed(c, e) {
		c.Status(204)
	}
}
func (h *Handler) EndSession(c *gin.Context, id int64) {
	if !noBody(c) {
		return
	}
	in := &api.RevokeAuthSessionRequest{Principal: principal(c), SessionId: &id}
	if in.Principal == nil {
		parts := strings.Fields(c.GetHeader("Authorization"))
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.Unauthorized(c, fmt.Errorf("missing bearer"))
			return
		}
		in.RefreshToken = parts[1]
	}
	_, e := h.service.Authentication.RevokeSessions(c, in)
	if !failed(c, e) {
		c.Status(204)
	}
}
