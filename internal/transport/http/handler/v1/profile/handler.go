package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/auth/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
)

type Handler struct {
	service *service.Auth
}

func NewHandler(service *service.Auth) *Handler {
	return &Handler{
		service: service,
	}
}

// DeleteProfile Swagger Documentation
//
//	@Summary		Delete Profile
//	@Description	Delete profile
//	@Tags			auth, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input		body		request_body.DeleteProfile	true	"Profile password"
//	@Success		200			{object}	response.MessageBody
//	@Failure		400			{object}	fail.Response
//	@Failure		401			{object}	fail.Response
//	@Failure		500			{object}	fail.Response
//	@Router			/v1/profile	[delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.DeleteProfile
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.DeleteProfile(c, &api.DeleteProfileRequest{
		Id:       payload.UserId.String(),
		Password: body.Password,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
