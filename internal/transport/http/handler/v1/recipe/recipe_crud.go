package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/response_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

// CreateRecipe Swagger Documentation
//
//	@Summary		Create recipe
//	@Description	Create recipe
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input		body		request_body.RecipeInput	true	"Input"
//	@Success		200			{object}	response_body.CreateRecipe
//	@Failure		400			{object}	fail.Response
//	@Failure		500			{object}	fail.Response
//	@Router			/v1/recipes	[post]
func (h *Handler) CreateRecipe(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.RecipeInput
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.CreateRecipe(c, request_body.RecipeInputRequest(body, payload.UserId))
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.CreateRecipe{
		RecipeId: res.RecipeId,
		Version:  res.Version,
	})
}

// GetRecipe Swagger Documentation
//
//	@Summary		Get recipe
//	@Description	Get recipe
//	@Tags			shopping-list
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id				path		string	true	"Recipe ID"
//	@Success		200						{object}	response_body.GetRecipeResponse
//	@Failure		400						{object}	fail.Response
//	@Failure		500						{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}	[get]
func (h *Handler) GetRecipe(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.GetRecipe
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.GetRecipe(c, &api.GetRecipeRequest{
		RecipeId:         c.Param(ParamRecipeId),
		UserId:           payload.UserId.String(),
		UserLanguage:     body.UserLanguage,
		Translate:        body.Translate,
		SubscriptionPlan: payload.SubscriptionPlan,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetRecipe(res))
}

// UpdateRecipe Swagger Documentation
//
//	@Summary		Delete recipe
//	@Description	Delete recipe
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id				path		string						true	"Recipe ID"
//	@Param			input					body		request_body.RecipeInput	true	"Input"
//	@Success		200						{object}	response_body.UpdateRecipe
//	@Failure		400						{object}	fail.Response
//	@Failure		500						{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}	[put]
func (h *Handler) UpdateRecipe(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	recipeId := c.Param(ParamRecipeId)
	var body request_body.RecipeInput
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}
	body.Id = &recipeId

	res, err := h.service.UpdateRecipe(c, request_body.RecipeInputRequest(body, payload.UserId))
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.UpdateRecipe{
		Version: res.Version,
	})
}

// DeleteRecipe Swagger Documentation
//
//	@Summary		Delete recipe
//	@Description	Delete recipe
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipe_id				path		string	true	"Recipe ID"
//	@Success		200						{object}	response.MessageBody
//	@Failure		400						{object}	fail.Response
//	@Failure		500						{object}	fail.Response
//	@Router			/v1/recipes/{recipe_id}	[delete]
func (h *Handler) DeleteRecipe(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.DeleteRecipe(c, &api.DeleteRecipeRequest{
		RecipeId: c.Param(ParamRecipeId),
		UserId:   payload.UserId.String(),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
