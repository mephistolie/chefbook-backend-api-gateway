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

// SetName Swagger Documentation
//
//	@Summary		Set name
//	@Description	Set profile name
//	@Tags			user, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.SetName	true	"Name"
//	@Success		200					{object}	response.MessageBody
//	@Failure		400					{object}	fail.Response
//	@Failure		401					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/profile/name	[put]
func (h *Handler) SetName(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SetName
	if err := c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.user.SetUserName(c, &api.SetUserNameRequest{
		UserId:    payload.UserId.String(),
		FirstName: body.FirstName,
		LastName:  body.LastName,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// SetDescription Swagger Documentation
//
//	@Summary		Set description
//	@Description	Set profile description
//	@Tags			user, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input					body		request_body.SetDescription	true	"Description"
//	@Success		200						{object}	response.MessageBody
//	@Failure		400						{object}	fail.Response
//	@Failure		401						{object}	fail.Response
//	@Failure		500						{object}	fail.Response
//	@Router			/v1/profile/description	[put]
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

// GenerateAvatarUploadLink Swagger Documentation
//
//	@Summary		Delete avatar
//	@Description	Delete avatar
//	@Tags			user, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	response.LinkBody
//	@Failure		400					{object}	fail.Response
//	@Failure		401					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/profile/avatar	[post]
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
		AvatarId: res.AvatarId,
		Link:     res.Link,
		FormData: res.FormData,
	})
}

// ConfirmAvatarUploading Swagger Documentation
//
//	@Summary		Confirm avatar uploading
//	@Description	Confirm avatar uploading
//	@Tags			user, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input				body		request_body.ConfirmAvatarUploading	true	"Avatar ID"
//	@Success		200					{object}	response.MessageBody
//	@Failure		400					{object}	fail.Response
//	@Failure		401					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/profile/avatar	[put]
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
		UserId:   payload.UserId.String(),
		AvatarId: body.AvatarId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// DeleteAvatar Swagger Documentation
//
//	@Summary		Delete avatar
//	@Description	Delete avatar
//	@Tags			user, profile
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	response.MessageBody
//	@Failure		400					{object}	fail.Response
//	@Failure		401					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/profile/avatar	[delete]
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
