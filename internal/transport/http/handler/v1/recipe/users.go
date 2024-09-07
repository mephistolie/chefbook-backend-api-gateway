package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

// RateRecipe Swagger Documentation
//
//	@Summary		Rate recipe
//	@Description	Rate recipe
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id						path		string	true	"Recipe ID"
//	@Success		200								{object}	response.MessageBody
//	@Failure		400								{object}	fail.Response
//	@Failure		500								{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}/rate	[post]
func (h *Handler) RateRecipe(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.RateRecipe
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}
	var score int32 = 0
	if body.Score != nil {
		score = *body.Score
	}

	res, err := h.service.RateRecipe(c, &api.RateRecipeRequest{
		RecipeId: c.Param(ParamRecipeId),
		UserId:   payload.UserId.String(),
		Score:    score,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// SaveRecipeToRecipeBook Swagger Documentation
//
//	@Summary		Save recipe
//	@Description	Save recipe to user's recipe book
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id						path		string	true	"Recipe ID"
//	@Success		200								{object}	response.MessageBody
//	@Failure		400								{object}	fail.Response
//	@Failure		500								{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}/book	[post]
func (h *Handler) SaveRecipeToRecipeBook(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.SaveRecipeToRecipeBook(c, &api.SaveRecipeToRecipeBookRequest{
		RecipeId: c.Param(ParamRecipeId),
		UserId:   payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// RemoveRecipeFromRecipeBook Swagger Documentation
//
//	@Summary		Remove recipe from recipe book
//	@Description	Remove recipe from recipe book
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id					path		string	true	"Recipe ID"
//	@Success		200							{object}	response.MessageBody
//	@Failure		400							{object}	fail.Response
//	@Failure		500							{object}	fail.Response
//	@Router			/v1/recipes{recipe_id}/book	[delete]
func (h *Handler) RemoveRecipeFromRecipeBook(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.RemoveRecipeFromRecipeBook(c, &api.RemoveRecipeFromRecipeBookRequest{
		RecipeId: c.Param(ParamRecipeId),
		UserId:   payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// SaveRecipeToFavourites Swagger Documentation
//
//	@Summary		Add recipe to favourite
//	@Description	Add recipe to favourite
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id							path		string	true	"Recipe ID"
//	@Success		200									{object}	response.MessageBody
//	@Failure		400									{object}	fail.Response
//	@Failure		500									{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}/favourites	[post]
func (h *Handler) SaveRecipeToFavourites(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.SaveRecipeToFavourites(c, &api.SaveRecipeToFavouritesRequest{
		RecipeId: c.Param(ParamRecipeId),
		UserId:   payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// RemoveRecipeFromFavourites Swagger Documentation
//
//	@Summary		Remove recipe from favourite
//	@Description	Remove recipe from favourite
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id							path		string	true	"Recipe ID"
//	@Success		200									{object}	response.MessageBody
//	@Failure		400									{object}	fail.Response
//	@Failure		500									{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}/favourite	[delete]
func (h *Handler) RemoveRecipeFromFavourites(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.RemoveRecipeFromFavourites(c, &api.RemoveRecipeFromFavouritesRequest{
		RecipeId: c.Param(ParamRecipeId),
		UserId:   payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// AddRecipeToCollection Swagger Documentation
//
//	@Summary		Add recipe to collection
//	@Description	Add recipe to collection
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id											path		string	true	"Recipe ID"
//	@Param			collection_id										path		string	true	"Collection ID"
//	@Success		200													{object}	response.MessageBody
//	@Failure		400													{object}	fail.Response
//	@Failure		500													{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}/collections/{collection_id}	[post]
func (h *Handler) AddRecipeToCollection(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.AddRecipeToCollection(c, &api.AddRecipeToCollectionRequest{
		RecipeId:     c.Param(ParamRecipeId),
		CollectionId: c.Param(ParamCollectionId),
		UserId:       payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// RemoveRecipeFromCollection Swagger Documentation
//
//	@Summary		Remove recipe from collection
//	@Description	Remove recipe from collection
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id											path		string	true	"Recipe ID"
//	@Param			collection_id										path		string	true	"Collection ID"
//	@Success		200													{object}	response.MessageBody
//	@Failure		400													{object}	fail.Response
//	@Failure		500													{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}/collections/{collection_id}	[delete]
func (h *Handler) RemoveRecipeFromCollection(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.RemoveRecipeFromCollection(c, &api.RemoveRecipeFromCollectionRequest{
		RecipeId:     c.Param(ParamRecipeId),
		CollectionId: c.Param(ParamCollectionId),
		UserId:       payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// SetRecipeCollections Swagger Documentation
//
//	@Summary		Set recipe collections
//	@Description	Set recipe collections
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id							path		string	true	"Recipe ID"
//	@Success		200									{object}	response.MessageBody
//	@Failure		400									{object}	fail.Response
//	@Failure		500									{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}/collections	[put]
func (h *Handler) SetRecipeCollections(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SetRecipeCollections
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}
	var collections []string
	if body.Collections != nil {
		collections = *body.Collections
	}

	res, err := h.service.SetRecipeCollections(c, &api.SetRecipeCollectionsRequest{
		RecipeId:      c.Param(ParamRecipeId),
		UserId:        payload.UserId.String(),
		CollectionIds: collections,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
