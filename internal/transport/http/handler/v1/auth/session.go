package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
)

// GetSessions Swagger Documentation
//
//	@Summary		Get Sessions
//	@Description	Get profile active sessions
//	@Tags			auth
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	response_body.Sessions
//	@Failure		400					{object}	fail.Response
//	@Failure		401					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/sessions	[get]
func (h *Handler) GetSessions(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetSessions(c, &api.GetSessionsRequest{Id: payload.UserId.String()})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.BySessions(res.Sessions, c.ClientIP()))
}

// EndSessions Swagger Documentation
//
//	@Summary		End Sessions
//	@Description	End selected sessions
//	@Tags			auth
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	response.MessageBody
//	@Failure		400					{object}	fail.Response
//	@Failure		401					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/auth/sessions	[delete]
func (h *Handler) EndSessions(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.Sessions
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.EndSessions(c, &api.EndSessionsRequest{
		Id:       payload.UserId.String(),
		Sessions: body,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
