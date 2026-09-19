package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
)

func (h *Handler) GetProfile(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	h.getProfileByIdOrUsername(c, payload.UserId, payload.UserId.String())
}

func (h *Handler) GetPublicProfile(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}
	h.getProfileByIdOrUsername(c, payload.UserId, c.Param(ParamProfileId))
}

func (h *Handler) getProfileByIdOrUsername(c *gin.Context, requesterId uuid.UUID, idOrName string) {
	profileId := ""
	username := ""

	if id, err := uuid.Parse(idOrName); err == nil {
		profileId = id.String()
	} else {
		username = idOrName
	}

	res, err := h.profile.GetProfile(c, &api.GetProfileRequest{
		ProfileId:       profileId,
		ProfileUsername: username,
		RequesterId:     requesterId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}
	response.Success(c, response_body.GetProfile(res))
}
