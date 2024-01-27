package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
)

// RequestPasswordReset Swagger Documentation
//
//	@Summary		Request Password Reset
//	@Description	Request password reset
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.RequestPasswordReset	true	"ProfileInfo identifier"
//	@Success		200					{object}	response.MessageBody
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/password	[get]
func (h *Handler) RequestPasswordReset(c *gin.Context) {
	var body request_body.RequestPasswordReset
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.RequestPasswordReset(c, &api.RequestPasswordResetRequest{
		Email:                    body.Email,
		Nickname:                 body.Nickname,
		ResetPasswordLinkPattern: h.routes.ResetPassword,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// ResetPassword Swagger Documentation
//
//	@Summary		Reset Password
//	@Description	Reset password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.ResetPassword	true	"Code and new password"
//	@Success		200					{object}	response.MessageBody
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/password	[post]
func (h *Handler) ResetPassword(c *gin.Context) {
	var body request_body.ResetPassword
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.ResetPassword(c, &api.ResetPasswordRequest{
		Id:          body.Id.String(),
		ResetCode:   body.ResetCode,
		NewPassword: body.NewPassword,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// ChangePassword Swagger Documentation
//
//	@Summary		Change Password
//	@Description	Change password
//	@Security		ApiKeyAuth
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.ChangePassword	true	"Code and new password"
//	@Success		200					{object}	response.MessageBody
//	@Failure		400					{object}	fail.Response
//	@Failure		401					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/password	[put]
func (h *Handler) ChangePassword(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.ChangePassword
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.ChangePassword(c, &api.ChangePasswordRequest{
		Id:          payload.UserId.String(),
		OldPassword: body.OldPassword,
		NewPassword: body.NewPassword,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
