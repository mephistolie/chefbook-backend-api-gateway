package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
)

// CheckUsernameAvailability Swagger Documentation
//
//	@Summary		Check Username Availability
//	@Description	Check profile username availability
//	@Tags			auth, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			username						path		string	true	"Username"
//	@Success		200								{object}	response_body.CheckUsername
//	@Failure		400								{object}	fail.Response
//	@Failure		401								{object}	fail.Response
//	@Failure		500								{object}	fail.Response
//	@Router			/v1/auth/username/{username} 	[get]
func (h *Handler) CheckUsernameAvailability(c *gin.Context) {
	res, err := h.service.CheckUsernameAvailability(c, &api.CheckUsernameAvailabilityRequest{
		Username: c.Param(ParamUsername),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.CheckUsername{Available: res.Available})
}

// SetUsername Swagger Documentation
//
//	@Summary		Set Username
//	@Description	Set profile username
//	@Tags			auth, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.Username	true	"Username"
//	@Success		200					{object}	response.MessageBody
//	@Failure		400					{object}	fail.Response
//	@Failure		401					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/username 	[post]
func (h *Handler) SetUsername(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.Username
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.SetUsername(c, &api.SetUsernameRequest{
		Id:       payload.UserId.String(),
		Username: body.Username,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
