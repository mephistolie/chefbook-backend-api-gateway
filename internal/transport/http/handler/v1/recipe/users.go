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

// SaveRecipe Swagger Documentation
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
//	@Router			/v1/recipes/{recipe_id}/save	[post]
func (h *Handler) SaveRecipe(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.SaveToRecipeBook(c, &api.SaveToRecipeBookRequest{
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
//	@Router			/v1/recipes{recipe_id}/save	[delete]
func (h *Handler) RemoveRecipeFromRecipeBook(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.RemoveFromRecipeBook(c, &api.RemoveFromRecipeBookRequest{
		RecipeId: c.Param(ParamRecipeId),
		UserId:   payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// AddRecipeToFavourite Swagger Documentation
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
//	@Router			/v1/recipes/{recipe_id}/favourite	[post]
func (h *Handler) AddRecipeToFavourite(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.SetRecipeFavouriteStatus(c, &api.SetRecipeFavouriteStatusRequest{
		RecipeId:  c.Param(ParamRecipeId),
		UserId:    payload.UserId.String(),
		Favourite: true,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// RemoveRecipeFromFavourite Swagger Documentation
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
func (h *Handler) RemoveRecipeFromFavourite(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.SetRecipeFavouriteStatus(c, &api.SetRecipeFavouriteStatusRequest{
		RecipeId:  c.Param(ParamRecipeId),
		UserId:    payload.UserId.String(),
		Favourite: false,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

// SetRecipeCategories Swagger Documentation
//
//	@Summary		Set recipe categories
//	@Description	Set recipe categories
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id							path		string	true	"Recipe ID"
//	@Success		200									{object}	response.MessageBody
//	@Failure		400									{object}	fail.Response
//	@Failure		500									{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}/categories	[put]
func (h *Handler) SetRecipeCategories(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.SetRecipeCategories
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}
	var categories []string
	if body.Categories != nil {
		categories = *body.Categories
	}

	res, err := h.service.SetRecipeCategories(c, &api.SetRecipeCategoriesRequest{
		RecipeId:    c.Param(ParamRecipeId),
		UserId:      payload.UserId.String(),
		CategoryIds: categories,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
