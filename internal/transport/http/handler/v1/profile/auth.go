package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"time"
)

func (h *Handler) GetProfileDeletionStatus(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.auth.GetProfileDeletionStatus(c, &api.GetProfileDeletionStatusRequest{
		ProfileId: payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	var deletionTimestamp *time.Time
	if res.DeletionTimestamp != nil {
		timestamp := res.DeletionTimestamp.AsTime()
		deletionTimestamp = &timestamp
	}

	response.Success(c, response_body.ProfileDeletionStatus{
		DeletionTimestamp: deletionTimestamp,
		Deleted:           res.Deleted,
	})
}

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

	res, err := h.auth.DeleteProfile(c, &api.DeleteProfileRequest{
		ProfileId:        payload.UserId.String(),
		Password:         body.Password,
		DeleteSharedData: body.WithSharedData,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Accepted(c, "/v1/profile/delete", response_body.DeleteProfile{
		DeletionTimestamp: res.DeletionTimestamp.AsTime(),
	})
}

func (h *Handler) CancelProfileDeletion(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.auth.CancelProfileDeletion(c, &api.CancelProfileDeletionRequest{
		ProfileId: payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
