package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-profile/api/proto/implementation/v1"
)

// GetProfile Swagger Documentation
//
//	@Summary		Get profile
//	@Description	Get profile
//	@Tags			profile, auth, user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			profile_id	path		int						true	"Shopping list ID"
//	@Success		200			{object}	response.MessageBody
//	@Failure		400			{object}	fail.Response
//	@Failure		401			{object}	fail.Response
//	@Failure		500			{object}	fail.Response
//	@Router			/v1/profile	[get]
func (h *Handler) GetProfile(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	h.getProfileByIdOrNickname(c, payload.UserId)
}

// getProfileById Swagger Documentation
//
//	@Summary		Get profile by ID
//	@Description	Get profile by ID or nickname
//	@Tags			profile, auth, user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			profile_id					path		string	true	"Profile ID or nickname"
//	@Success		200							{object}	response_body.Profile
//	@Failure		400							{object}	fail.Response
//	@Failure		401							{object}	fail.Response
//	@Failure		500							{object}	fail.Response
//	@Router			/v1/profiles/{profile_id}	[get]
func (h *Handler) getProfileByIdOrNickname(c *gin.Context, requesterId uuid.UUID) {
	profileId := ""
	nickname := ""

	idOrName := c.Param(ParamProfileId)
	if id, err := uuid.Parse(idOrName); err == nil {
		profileId = id.String()
	} else {
		nickname = idOrName
	}

	res, err := h.profile.GetProfile(c, &api.GetProfileRequest{
		ProfileId:       profileId,
		ProfileNickname: nickname,
		RequesterId:     requesterId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}
	response.Success(c, response_body.GetProfile(res))
}
