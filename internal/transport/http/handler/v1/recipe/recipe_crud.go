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

	response.Created(c, "/v1/recipes/"+url.PathEscape(res.RecipeId), response_body.CreateRecipe{
		RecipeId: res.RecipeId,
		Version:  res.Version,
	})
}

func (h *Handler) GetRecipe(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var languagePtr *string
	if language := c.Query(queryLanguage); len(language) > 0 {
		languagePtr = &language
	}
	var translatorIdPtr *string
	if translatorId := c.Query(queryTranslatorId); len(translatorId) > 0 {
		translatorIdPtr = &translatorId
	}

	res, err := h.service.GetRecipe(c, &api.GetRecipeRequest{
		RecipeId:         c.Param(ParamRecipeId),
		UserId:           payload.UserId.String(),
		Language:         languagePtr,
		TranslatorId:     translatorIdPtr,
		Translate:        c.Query(queryTranslated) == "true",
		SubscriptionPlan: payload.SubscriptionPlan,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Success(c, response_body.GetRecipe(res))
}

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
