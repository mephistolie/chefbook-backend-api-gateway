package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/profile/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
	"net/http"
)

func (h *Handler) SetDisplayName(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SetDisplayName
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.user.SetUserDisplayName(c, &api.SetUserDisplayNameRequest{
		UserId:      payload.UserId.String(),
		DisplayName: body.DisplayName,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

func (h *Handler) SetDescription(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SetDescription
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.user.SetUserDescription(c, &api.SetUserDescriptionRequest{
		UserId:      payload.UserId.String(),
		Description: body.Description,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

func (h *Handler) GenerateAvatarUploadLink(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.user.GenerateUserAvatarUploadLink(c, &api.GenerateUserAvatarUploadLinkRequest{
		UserId: payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	c.JSON(http.StatusOK, response_body.GenerateAvatarUploadLink{
		PictureLink: res.AvatarLink,
		UploadLink:  res.UploadLink,
		FormData:    response.NonNilStringMap(res.FormData),
		MaxSize:     res.MaxSize,
	})
}

func (h *Handler) ConfirmAvatarUploading(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.ConfirmAvatarUploading
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.user.ConfirmUserAvatarUploading(c, &api.ConfirmUserAvatarUploadingRequest{
		UserId:     payload.UserId.String(),
		AvatarLink: body.AvatarLink,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

func (h *Handler) DeleteAvatar(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.user.DeleteUserAvatar(c, &api.DeleteUserAvatarRequest{
		UserId: payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
