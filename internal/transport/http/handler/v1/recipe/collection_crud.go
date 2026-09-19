package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
	"net/url"
)

const (
	queryUserId = "user_id"
)

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

	response.Created(c, "/v1/collections/"+url.PathEscape(res.CollectionId), response_body.AddCollection{Id: res.CollectionId})
}

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
