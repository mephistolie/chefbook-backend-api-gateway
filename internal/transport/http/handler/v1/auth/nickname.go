package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
)

// CheckNicknameAvailability Swagger Documentation
//
//	@Summary		Check Nickname Availability
//	@Description	Check profile nickname availability
//	@Tags			auth, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.Nickname	true	"Nickname"
//	@Success		200					{object}	response_body.CheckNickname
//	@Failure		400					{object}	fail.Response
//	@Failure		401					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/nickname 	[get]
func (h *Handler) CheckNicknameAvailability(c *gin.Context) {
	var body request_body.Nickname
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.CheckNicknameAvailability(c, &api.CheckNicknameAvailabilityRequest{Nickname: body.Nickname})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.CheckNickname{Available: res.Available})
}

// SetNickname Swagger Documentation
//
//	@Summary		Set Nickname
//	@Description	Set profile nickname
//	@Tags			auth, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.Nickname	true	"Nickname"
//	@Success		200					{object}	response.MessageBody
//	@Failure		400					{object}	fail.Response
//	@Failure		401					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/nickname 	[post]
func (h *Handler) SetNickname(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.Nickname
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.SetNickname(c, &api.SetNicknameRequest{
		Id:       payload.UserId.String(),
		Nickname: body.Nickname,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
