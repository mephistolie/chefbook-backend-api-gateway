package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"strconv"
)

func (h *Handler) ChangePassword(c *gin.Context) {
	var b struct {
		Password string `json:"newPassword"`
	}
	if !bind(c, "SetPasswordRequest", &b) {
		return
	}
	_, e := h.service.Authentication.SetPassword(c, &api.SetAccountPasswordRequest{Action: sensitive(c), NewPassword: b.Password})
	if !failed(c, e) {
		c.Status(204)
	}
}
func (h *Handler) RequestPasswordReset(c *gin.Context) {
	var b struct {
		Email string `json:"email"`
	}
	if !bind(c, "RequestPasswordResetRequest", &b) {
		return
	}
	_, e := h.service.Authentication.StartPasswordReset(c, &api.StartPasswordResetRequest{Email: b.Email})
	if !failed(c, e) {
		c.Status(202)
	}
}
func (h *Handler) ResetPassword(c *gin.Context) {
	var b struct {
		Token    string `json:"token"`
		Password string `json:"newPassword"`
	}
	if !bind(c, "ConfirmPasswordResetRequest", &b) {
		return
	}
	_, e := h.service.Authentication.ConfirmPasswordReset(c, &api.ConfirmPasswordResetRequest{Token: b.Token, NewPassword: b.Password})
	if !failed(c, e) {
		c.Status(204)
	}
}
func (h *Handler) SetUsername(c *gin.Context) {
	var b struct {
		Username string `json:"username"`
	}
	if !bind(c, "SetUsernameRequest", &b) {
		return
	}
	_, e := h.service.Authentication.SetUsername(c, &api.SetAccountUsernameRequest{Principal: principal(c), Username: b.Username})
	if !failed(c, e) {
		c.Status(204)
	}
}
func (h *Handler) CheckUsernameAvailability(c *gin.Context) {
	s, e := h.service.Authentication.UsernameAvailability(c, &api.AuthUsernameRequest{Username: c.Param("username")})
	if !failed(c, e) {
		c.JSON(200, gin.H{"available": s.Available})
	}
}
func (h *Handler) RequestEmailVerification(c *gin.Context) {
	var b struct {
		Email string `json:"email"`
	}
	if !bind(c, "RequestEmailChangeRequest", &b) {
		return
	}
	_, e := h.service.Authentication.ChangeEmail(c, &api.ChangeAccountEmailRequest{Action: sensitive(c), Email: b.Email})
	if !failed(c, e) {
		c.Status(202)
	}
}
func (h *Handler) ConfirmEmailChange(c *gin.Context) {
	var b struct {
		Token string `json:"token"`
	}
	if !bind(c, "ConfirmEmailChangeRequest", &b) {
		return
	}
	s, e := h.service.Authentication.ConfirmEmail(c, &api.ConfirmAccountEmailRequest{Token: b.Token})
	if !failed(c, e) {
		c.JSON(200, gin.H{"status": s.Status})
	}
}
func deletionJSON(s *api.AccountDeletion) gin.H {
	return gin.H{"deletionTimestamp": timestamp(s.DeletionTimestamp), "deleteSharedData": s.DeleteSharedData}
}
func (h *Handler) DeleteProfile(c *gin.Context) {
	var b struct {
		DeleteSharedData bool `json:"deleteSharedData"`
	}
	if !bind(c, "RequestAccountDeletionRequest", &b) {
		return
	}
	s, e := h.service.Authentication.RequestDeletion(c, &api.RequestAccountDeletionRequest{Action: sensitive(c), DeleteSharedData: b.DeleteSharedData})
	if !failed(c, e) {
		c.Header("Location", "/v1/account/deletion")
		c.JSON(202, deletionJSON(s))
	}
}
func (h *Handler) UpdateAccountDeletion(c *gin.Context) {
	var b struct {
		DeleteSharedData bool `json:"deleteSharedData"`
	}
	if !bind(c, "UpdateAccountDeletionRequest", &b) {
		return
	}
	s, e := h.service.Authentication.PatchDeletion(c, &api.PatchAccountDeletionRequest{Principal: principal(c), DeleteSharedData: b.DeleteSharedData})
	if !failed(c, e) {
		c.JSON(200, deletionJSON(s))
	}
}
func (h *Handler) CancelProfileDeletion(c *gin.Context) {
	if !noBody(c) {
		return
	}
	_, e := h.service.Authentication.CancelDeletion(c, principal(c))
	if !failed(c, e) {
		c.Status(204)
	}
}
func (h *Handler) GetIdentities(c *gin.Context) {
	s, e := h.service.Authentication.GetIdentities(c, principal(c))
	if failed(c, e) {
		return
	}
	ids := []gin.H{}
	for _, v := range s.Identities {
		var id any = v.Subject
		if v.Provider == "vk" {
			n, err := strconv.ParseInt(v.Subject, 10, 64)
			if err != nil {
				invalidProviderResult(c)
				return
			}
			id = n
		}
		ids = append(ids, gin.H{"provider": v.Provider, "id": id})
	}
	c.JSON(200, gin.H{"identities": ids})
}
func (h *Handler) linkIdentity(c *gin.Context, provider, model string) {
	var b struct {
		Code  string `json:"code"`
		State string `json:"state"`
	}
	if !bind(c, model, &b) {
		return
	}
	s, e := h.service.Authentication.LinkIdentity(c, &api.LinkIdentityRequest{Action: sensitive(c), Provider: provider, Code: b.Code, State: b.State})
	if failed(c, e) {
		return
	}
	if s.Created {
		c.Header("Location", "/v1/account/identities/"+provider)
		c.Status(201)
	} else {
		c.Status(204)
	}
}
func (h *Handler) unlinkIdentity(c *gin.Context, provider string) {
	if !noBody(c) {
		return
	}
	_, e := h.service.Authentication.UnlinkIdentity(c, &api.UnlinkIdentityRequest{Action: sensitive(c), Provider: provider})
	if !failed(c, e) {
		c.Status(204)
	}
}
func (h *Handler) ConnectGoogle(c *gin.Context) {
	h.linkIdentity(c, "google", "LinkGoogleIdentityRequest")
}
func (h *Handler) ConnectVk(c *gin.Context)              { h.linkIdentity(c, "vk", "LinkVkIdentityRequest") }
func (h *Handler) DeleteGoogleConnection(c *gin.Context) { h.unlinkIdentity(c, "google") }
func (h *Handler) DeleteVkConnection(c *gin.Context)     { h.unlinkIdentity(c, "vk") }
func (h *Handler) prepareOAuth(c *gin.Context, provider string) {
	var b struct {
		RedirectURI string `json:"redirectUri"`
	}
	if !bind(c, "CreateAuthorizationUrlRequest", &b) {
		return
	}
	s, e := h.service.Authentication.PrepareIdentityOAuth(c, &api.PrepareIdentityOAuthRequest{Principal: principal(c), Provider: provider, RedirectUri: b.RedirectURI})
	if !failed(c, e) {
		c.JSON(201, gin.H{"url": s.Url, "expirationTimestamp": timestamp(s.ExpirationTimestamp)})
	}
}
func (h *Handler) RequestGoogleOAuth(c *gin.Context) { h.prepareOAuth(c, "google") }
func (h *Handler) RequestVkOAuth(c *gin.Context)     { h.prepareOAuth(c, "vk") }
func (h *Handler) GetTotp(c *gin.Context) {
	s, e := h.service.Authentication.GetTotp(c, principal(c))
	if !failed(c, e) {
		c.JSON(200, gin.H{"enabled": s.Enabled, "activationTimestamp": timestamp(s.ActivationTimestamp)})
	}
}
func (h *Handler) CreateTotp(c *gin.Context) {
	if !noBody(c) {
		return
	}
	s, e := h.service.Authentication.StartTotp(c, sensitive(c))
	if !failed(c, e) {
		c.Header("Location", "/v1/account/totp")
		c.JSON(201, gin.H{"secret": s.Secret, "otpauthUri": s.OtpauthUri, "expirationTimestamp": timestamp(s.ExpirationTimestamp)})
	}
}
func backupJSON(s *api.BackupCodes) gin.H {
	return gin.H{"backupCodes": emptyStrings(s.Codes), "generationTimestamp": timestamp(s.GenerationTimestamp)}
}
func (h *Handler) ConfirmTotp(c *gin.Context) {
	var b struct {
		Code string `json:"code"`
	}
	if !bind(c, "ConfirmTotpActivationRequest", &b) {
		return
	}
	s, e := h.service.Authentication.ConfirmTotp(c, &api.ConfirmTotpRequest{Principal: principal(c), Code: b.Code})
	if !failed(c, e) {
		c.JSON(200, backupJSON(s))
	}
}
func (h *Handler) DeleteTotp(c *gin.Context) {
	if !noBody(c) {
		return
	}
	_, e := h.service.Authentication.DeleteTotp(c, sensitive(c))
	if !failed(c, e) {
		c.Status(204)
	}
}
func (h *Handler) GetBackupCodes(c *gin.Context) {
	s, e := h.service.Authentication.GetBackupCodes(c, principal(c))
	if !failed(c, e) {
		c.JSON(200, gin.H{"remainingCount": s.RemainingCount, "generationTimestamp": timestamp(s.GenerationTimestamp)})
	}
}
func (h *Handler) CreateBackupCodes(c *gin.Context) {
	if !noBody(c) {
		return
	}
	s, e := h.service.Authentication.RotateBackupCodes(c, sensitive(c))
	if !failed(c, e) {
		c.Header("Location", "/v1/account/backup-codes")
		c.JSON(201, backupJSON(s))
	}
}
func passkeyJSON(s *api.AuthPasskey) gin.H {
	return gin.H{"id": s.Id, "name": s.Name, "creationTimestamp": timestamp(s.CreationTimestamp), "lastUseTimestamp": timestamp(s.LastUseTimestamp), "transports": emptyStrings(s.Transports), "backupEligible": s.BackupEligible, "backupState": s.BackupState}
}
func (h *Handler) GetPasskeys(c *gin.Context) {
	s, e := h.service.Authentication.GetPasskeys(c, principal(c))
	if failed(c, e) {
		return
	}
	keys := []gin.H{}
	for _, key := range s.Passkeys {
		keys = append(keys, passkeyJSON(key))
	}
	c.JSON(200, gin.H{"passkeys": keys})
}
func (h *Handler) CreatePasskeyRegistration(c *gin.Context) {
	if !noBody(c) {
		return
	}
	s, e := h.service.Authentication.StartPasskeyRegistration(c, sensitive(c))
	if !failed(c, e) {
		if !json.Valid(s.PublicKeyOptions) {
			invalidProviderResult(c)
			return
		}
		c.Header("Location", "/v1/account/passkeys/requests/"+s.RequestId)
		c.JSON(201, gin.H{"requestId": s.RequestId, "expirationTimestamp": timestamp(s.ExpirationTimestamp), "publicKey": json.RawMessage(s.PublicKeyOptions)})
	}
}
func (h *Handler) CreatePasskey(c *gin.Context) {
	var b struct {
		RequestID  string          `json:"requestId"`
		Name       string          `json:"name"`
		Credential json.RawMessage `json:"credential"`
	}
	if !bind(c, "ConfirmPasskeyRegistrationRequest", &b) {
		return
	}
	s, e := h.service.Authentication.RegisterPasskey(c, &api.RegisterAuthPasskeyRequest{Principal: principal(c), RequestId: b.RequestID, Name: b.Name, Credential: b.Credential})
	if !failed(c, e) {
		c.Header("Location", "/v1/account/passkeys/"+s.Id)
		c.JSON(201, passkeyJSON(s))
	}
}
func (h *Handler) RenamePasskey(c *gin.Context) {
	var b struct {
		Name string `json:"name"`
	}
	if !bind(c, "RenamePasskeyRequest", &b) {
		return
	}
	s, e := h.service.Authentication.RenamePasskey(c, &api.RenameAuthPasskeyRequest{Principal: principal(c), Id: c.Param("id"), Name: b.Name})
	if !failed(c, e) {
		c.JSON(200, passkeyJSON(s))
	}
}
func (h *Handler) DeletePasskey(c *gin.Context) {
	if !noBody(c) {
		return
	}
	_, e := h.service.Authentication.DeletePasskey(c, &api.DeleteAuthPasskeyRequest{Action: sensitive(c), Id: c.Param("id")})
	if !failed(c, e) {
		c.Status(204)
	}
}
