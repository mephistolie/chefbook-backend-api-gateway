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
	queryLanguage = "language"
)

// GetRecipes Swagger Documentation
//
//	@Summary		Get recipes
//	@Description	Get recipes by query
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input		body		request_body.GetRecipes	true	"Query"
//	@Success		200			{object}	response_body.GetRecipesResponse
//	@Failure		400			{object}	fail.Response
//	@Failure		500			{object}	fail.Response
//	@Router			/v1/recipes	[get]
func (h *Handler) GetRecipes(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.GetRecipes
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	res, err := h.service.GetRecipes(c, request_body.GetRecipesRequest(body, payload.UserId))
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetRecipes(res))
}

// GetRandomRecipe Swagger Documentation
//
//	@Summary		Get random recipe
//	@Description	Get random recipe
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			recipeLanguage		query		[]string	false	"Recipe language codes"
//	@Param			userLanguage		query		string		false	"User language code"
//	@Success		200					{object}	response_body.GetRecipeResponse
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/recipes/random	[get]
func (h *Handler) GetRandomRecipe(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var languagePtr *string
	language := c.Query(queryLanguage)
	if len(language) > 0 {
		languagePtr = &language
	}

	res, err := h.service.GetRandomRecipe(c, &api.GetRandomRecipeRequest{
		UserId:          payload.UserId.String(),
		RecipeLanguages: c.QueryArray(queryLanguage),
		UserLanguage:    languagePtr,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetRecipe(res))
}

// GetRecipeBook Swagger Documentation
//
//	@Summary		Get recipe book
//	@Description	Get user recipe book
//	@Tags			recipe
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			user_language		query		string	false	"User language code"
//	@Success		200					{object}	response_body.GetRecipeBookResponse
//	@Failure		400					{object}	fail.Response
//	@Failure		500					{object}	fail.Response
//	@Router			/v1/recipes/book	[get]
func (h *Handler) GetRecipeBook(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var languagePtr *string
	language := c.Query(queryLanguage)
	if len(language) > 0 {
		languagePtr = &language
	}

	res, err := h.service.GetRecipeBook(c, &api.GetRecipeBookRequest{
		UserId:       payload.UserId.String(),
		UserLanguage: languagePtr,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetRecipeBook(res))
}
