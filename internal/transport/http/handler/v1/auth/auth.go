package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"time"
)

// SignUp Swagger Documentation
//
//	@Summary		Sign Up
//	@Description	Create new profile
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.SignUp	true	"Credentials"
//	@Success		200					{object}	response_body.SignUp
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/sign-up	[post]
func (h *Handler) SignUp(c *gin.Context) {
	var body request_body.SignUp
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	id := ""
	if body.Id != nil {
		id = body.Id.String()
	}

	res, err := h.service.SignUp(c, &api.SignUpRequest{
		Id:                    id,
		Email:                 body.Email,
		Password:              body.Password,
		ActivationLinkPattern: h.routes.ActivateProfile,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	message := ""
	if !res.Activated {
		message = "activation link has been sent"
	}

	response.Success(c, response_body.SignUp{
		Id:        res.Id,
		Activated: res.Activated,
		Message:   message,
	})
}

// ActivateProfile Swagger Documentation
//
//	@Summary		Activate Profile
//	@Description	Activate profile
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user_id				query		string	true	"User ID"
//	@Param			code				query		string	true	"Activation code"
//	@Success		200					{object}	response.MessageBody
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/activate 	[get]
func (h *Handler) ActivateProfile(c *gin.Context) {
	res, err := h.service.ActivateProfile(c, &api.ActivateProfileRequest{
		Id:             c.Query("user_id"),
		ActivationCode: c.Query("code"),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// SignIn Swagger Documentation
//
//	@Summary		Sign In
//	@Description	Sign in to profile
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.SignIn	true	"Credentials"
//	@Success		200					{object}	response_body.Tokens
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/sign-in	[post]
func (h *Handler) SignIn(c *gin.Context) {
	var body request_body.SignIn
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.SignIn(c, &api.SignInRequest{
		Email:     body.Email,
		Nickname:  body.Nickname,
		Password:  body.Password,
		Ip:        c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	var profileDeletionTimestamp *time.Time
	if res.ProfileDeletionTimestamp != nil {
		timestamp := res.ProfileDeletionTimestamp.AsTime()
		profileDeletionTimestamp = &timestamp
	}

	response.Success(c, response_body.Tokens{
		Access:            res.AccessToken,
		Refresh:           res.RefreshToken,
		ExpiresAt:         res.ExpirationTimestamp.AsTime(),
		ProfileDeletingAt: profileDeletionTimestamp,
	})
}

// RefreshSession Swagger Documentation
//
//	@Summary		Refresh Session
//	@Description	Refresh session to get new tokens pair
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.RefreshToken	true	"Refresh token"
//	@Success		200					{object}	response_body.Tokens
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/refresh	[post]
func (h *Handler) RefreshSession(c *gin.Context) {
	var body request_body.RefreshToken
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.RefreshSession(c, &api.RefreshSessionRequest{
		RefreshToken: body.RefreshToken,
		Ip:           c.ClientIP(),
		UserAgent:    c.Request.UserAgent(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	var profileDeletionTimestamp *time.Time
	if res.ProfileDeletionTimestamp != nil {
		timestamp := res.ProfileDeletionTimestamp.AsTime()
		profileDeletionTimestamp = &timestamp
	}

	response.Success(c, response_body.Tokens{
		Access:            res.AccessToken,
		Refresh:           res.RefreshToken,
		ExpiresAt:         res.ExpirationTimestamp.AsTime(),
		ProfileDeletingAt: profileDeletionTimestamp,
	})
}

// SignOut Swagger Documentation
//
//	@Summary		Sign Out
//	@Description	Sign out and remove session
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.RefreshToken	true	"Refresh token"
//	@Success		200					{object}	response.MessageBody
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/sign-out 	[post]
func (h *Handler) SignOut(c *gin.Context) {
	var body request_body.RefreshToken
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.SignOut(c, &api.SignOutRequest{RefreshToken: body.RefreshToken})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
