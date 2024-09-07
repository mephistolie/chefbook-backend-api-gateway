package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

const (
	queryUserId = "user_id"
)

// GetCollections Swagger Documentation
//
//	@Summary		Get user collections
//	@Description	Get user recipe collections
//	@Tags			collection, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			user_id			query		string	false	"User ID"
//	@Success		200				{object}	[]response_body.Collection
//	@Failure		400				{object}	fail.Response
//	@Failure		500				{object}	fail.Response
//	@Router			/v1/collections	[get]
func (h *Handler) GetCollections(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	userId := payload.UserId.String()
	if specifiedUserId := c.Query(queryUserId); len(specifiedUserId) > 0 {
		userId = specifiedUserId
	}

	res, err := h.service.GetCollections(c, &api.GetCollectionsRequest{
		UserId:      userId,
		RequesterId: payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.NewGetCollections(res))
}

// AddCollection Swagger Documentation
//
//	@Summary		Add collection
//	@Description	Add recipes collection
//	@Tags			collection, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input			body		request_body.AddCollection	true	"Input"
//	@Success		200				{object}	response_body.AddCollection
//	@Failure		400				{object}	fail.Response
//	@Failure		500				{object}	fail.Response
//	@Router			/v1/collections	[post]
func (h *Handler) AddCollection(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.AddCollection
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.CreateCollection(c, &api.CreateCollectionRequest{
		UserId:       payload.UserId.String(),
		CollectionId: body.Id,
		Name:         body.Name,
		Visibility:   body.Visibility,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.AddCollection{Id: res.CollectionId})
}

// GetCollection Swagger Documentation
//
//	@Summary		Get collection
//	@Description	Get recipes collection
//	@Tags			collection, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			collection_id					path		string	true	"Collection ID"
//	@Success		200								{object}	response_body.GetCollection
//	@Failure		400								{object}	fail.Response
//	@Failure		500								{object}	fail.Response
//	@Router			/v1/collections/{collection_id}	[get]
func (h *Handler) GetCollection(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.GetCollection(c, &api.GetCollectionRequest{
		CollectionId: c.Param(ParamCollectionId),
		UserId:       payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.NewGetCollection(res))
}

// UpdateCollection Swagger Documentation
//
//	@Summary		Update collection
//	@Description	Update recipes collection
//	@Tags			collection, recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			collection_id					path		int								true	"Collection ID"
//	@Param			input							body		request_body.UpdateCollection	true	"Input"
//	@Success		200								{object}	response.MessageBody
//	@Failure		400								{object}	fail.Response
//	@Failure		500								{object}	fail.Response
//	@Router			/v1/collections/{collection_id}	[put]
func (h *Handler) UpdateCollection(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.UpdateCollection
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.UpdateCollection(c, &api.UpdateCollectionRequest{
		UserId:       payload.UserId.String(),
		CollectionId: c.Param(ParamCollectionId),
		Name:         body.Name,
		Visibility:   body.Visibility,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// DeleteCollection Swagger Documentation
//
//	@Summary		Delete collection
//	@Description	Delete recipes collection
//	@Tags			collection
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			collection_id					path		int	true	"Collection ID"
//	@Success		200								{object}	response.MessageBody
//	@Failure		400								{object}	fail.Response
//	@Failure		500								{object}	fail.Response
//	@Router			/v1/collections/{collection_id}	[delete]
func (h *Handler) DeleteCollection(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.DeleteCollection(c, &api.DeleteCollectionRequest{
		CollectionId: c.Param(ParamCollectionId),
		UserId:       payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
