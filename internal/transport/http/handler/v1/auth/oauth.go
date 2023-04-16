package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
)

// RequestGoogleOAuth Swagger Documentation
//
//	@Summary		Request Google OAuth
//	@Description	Request Google OAuth link
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200						{object}	response_body.Link
//	@Failure		500						{object}	fail.Response
//	@Router			/v1/auth/google/request	[get]
func (h *Handler) RequestGoogleOAuth(c *gin.Context) {
	res, err := h.service.RequestGoogleOAuth(c, &api.RequestGoogleOAuthRequest{
		RedirectUrl: h.routes.SignInGoogle,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Link(c, res.Link)
}

// SignInGoogle Swagger Documentation
//
//	@Summary		Sign In Google
//	@Description	Sign in to profile via Google Account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input			body		request_body.OAuthCode	true	"Code"
//	@Success		200				{object}	response_body.Tokens
//	@Failure		400				{object}	fail.Response
//	@Failure		500				{object}	fail.Response
//	@Router			/v1/auth/google	[post]
func (h *Handler) SignInGoogle(c *gin.Context) {
	var body request_body.OAuthCode
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response_body.InvalidBody)
		return
	}

	res, err := h.service.SignInGoogle(c, &api.SignInGoogleRequest{
		Code:        body.Code,
		State:       body.State,
		RedirectUrl: h.routes.SignInGoogle,
		Ip:          c.ClientIP(),
		UserAgent:   c.Request.UserAgent(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.Tokens{
		Access:    res.AccessToken,
		Refresh:   res.RefreshToken,
		ExpiresAt: res.ExpirationTimestamp.AsTime(),
	})
}

// ConnectGoogle Swagger Documentation
//
//	@Summary		Connect Google
//	@Description	Connect Google to existing profile
//	@Tags			auth
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input			body		request_body.OAuthCode	true	"Credentials"
//	@Success		200				{object}	response_body.Message
//	@Failure		400				{object}	fail.Response
//	@Failure		401				{object}	fail.Response
//	@Failure		500				{object}	fail.Response
//	@Router			/v1/auth/google	[put]
func (h *Handler) ConnectGoogle(c *gin.Context) {
	payload, err := request.GetUserPayload(c)
	if err != nil {
		response.Unknown(c, err)
		return
	}

	var body request_body.OAuthCode
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response_body.InvalidBody)
		return
	}

	res, err := h.service.ConnectGoogle(c, &api.ConnectGoogleRequest{
		Id:          payload.UserId.String(),
		Code:        body.Code,
		State:       body.State,
		RedirectUrl: h.routes.SignInGoogle,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// DeleteGoogleConnection Swagger Documentation
//
//	@Summary		Delete Google Connection
//	@Description	Delete Google connection for profile
//	@Tags			auth
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	response_body.Message
//	@Failure		400				{object}	fail.Response
//	@Failure		401				{object}	fail.Response
//	@Failure		500				{object}	fail.Response
//	@Router			/v1/auth/google	[delete]
func (h *Handler) DeleteGoogleConnection(c *gin.Context) {
	payload, err := request.GetUserPayload(c)
	if err != nil {
		response.Unknown(c, err)
		return
	}

	res, err := h.service.DeleteGoogleConnection(c, &api.DeleteGoogleConnectionRequest{Id: payload.UserId.String()})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// RequestVkOAuth Swagger Documentation
//
//	@Summary		Request VK OAuth
//	@Description	Request VK OAuth link
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			display				query		string	false	"Display"
//	@Param			response_type		query		string	false	"Response type"
//	@Success		200					{object}	response_body.Link
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/vk/request	[get]
func (h *Handler) RequestVkOAuth(c *gin.Context) {
	res, err := h.service.RequestVkOAuth(c, &api.RequestVkOAuthRequest{
		Display:      c.Query("display"),
		ResponseType: c.Query("response_type"),
		RedirectUri:  h.routes.SignInVk,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Link(c, res.Link)
}

// SignInVk Swagger Documentation
//
//	@Summary		Sign In VK
//	@Description	Sign in to profile via VK Account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input		body		request_body.OAuthCode	true	"Credentials"
//	@Success		200			{object}	response_body.Tokens
//	@Failure		400			{object}	fail.Response
//	@Failure		500			{object}	fail.Response
//	@Router			/v1/auth/vk	[post]
func (h *Handler) SignInVk(c *gin.Context) {
	var body request_body.OAuthCode
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response_body.InvalidBody)
		return
	}

	res, err := h.service.SignInVk(c, &api.SignInVkRequest{
		Code:        body.Code,
		State:       body.State,
		RedirectUri: h.routes.SignInVk,
		Ip:          c.ClientIP(),
		UserAgent:   c.Request.UserAgent(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.Tokens{
		Access:    res.AccessToken,
		Refresh:   res.RefreshToken,
		ExpiresAt: res.ExpirationTimestamp.AsTime(),
	})
}

// ConnectVk Swagger Documentation
//
//	@Summary		Connect VK
//	@Description	Connect VK to existing profile
//	@Tags			auth
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input		body		request_body.OAuthCode	true	"Credentials"
//	@Success		200			{object}	response_body.Message
//	@Failure		400			{object}	fail.Response
//	@Failure		401			{object}	fail.Response
//	@Failure		500			{object}	fail.Response
//	@Router			/v1/auth/vk	[put]
func (h *Handler) ConnectVk(c *gin.Context) {
	payload, err := request.GetUserPayload(c)
	if err != nil {
		response.Unknown(c, err)
		return
	}

	var body request_body.OAuthCode
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response_body.InvalidBody)
		return
	}

	res, err := h.service.ConnectVk(c, &api.ConnectVkRequest{
		Id:          payload.UserId.String(),
		Code:        body.Code,
		State:       body.State,
		RedirectUri: h.routes.SignInVk,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// DeleteVkConnection Swagger Documentation
//
//	@Summary		Delete VK Connection
//	@Description	Delete VK connection for profile
//	@Tags			auth
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	response_body.Message
//	@Failure		400			{object}	fail.Response
//	@Failure		401			{object}	fail.Response
//	@Failure		500			{object}	fail.Response
//	@Router			/v1/auth/vk	[delete]
func (h *Handler) DeleteVkConnection(c *gin.Context) {
	payload, err := request.GetUserPayload(c)
	if err != nil {
		response.Unknown(c, err)
		return
	}

	res, err := h.service.DeleteVkConnection(c, &api.DeleteVkConnectionRequest{Id: payload.UserId.String()})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
