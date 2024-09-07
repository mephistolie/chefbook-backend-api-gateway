package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

// SaveCollectionToRecipeBook Swagger Documentation
//
//	@Summary		Save collection to recipe book
//	@Description	Save collection to recipe book
//	@Tags			collection
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			collection_id							path		int	true	"Collection ID"
//	@Success		200										{object}	response.MessageBody
//	@Failure		400										{object}	fail.Response
//	@Failure		500										{object}	fail.Response
//	@Router			/v1/collections/{collection_id}/save	[post]
func (h *Handler) SaveCollectionToRecipeBook(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SaveCollectionToRecipeBook
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.SaveCollectionToRecipeBook(c, &api.SaveCollectionToRecipeBookRequest{
		CollectionId:   c.Param(ParamCollectionId),
		UserId:         payload.UserId.String(),
		ContributorKey: body.ContributorKey,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// RemoveCollectionFromRecipeBook Swagger Documentation
//
//	@Summary		Remove collection from recipe book
//	@Description	Remove collection from recipe book
//	@Tags			collection
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			collection_id							path		int	true	"Collection ID"
//	@Success		200										{object}	response.MessageBody
//	@Failure		400										{object}	fail.Response
//	@Failure		500										{object}	fail.Response
//	@Router			/v1/collections/{collection_id}/save	[delete]
func (h *Handler) RemoveCollectionFromRecipeBook(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.RemoveCollectionFromRecipeBook(c, &api.RemoveCollectionFromRecipeBookRequest{
		CollectionId: c.Param(ParamCollectionId),
		UserId:       payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
