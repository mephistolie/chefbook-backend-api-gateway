package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/request_body"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

const (
	queryTranslatorId = "translator_id"
	queryTranslated   = "translated"
)

func (h *Handler) TranslateRecipe(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	var body request_body.TranslateRecipe
	if err = c.BindJSON(&body); err != nil {
		response.Fail(c, response.InvalidBody)
		return
	}

	ingredients := map[string]*api.IngredientTranslation{}
	for id, translation := range body.Ingredients {
		ingredients[id] = &api.IngredientTranslation{
			Text: translation.Text,
			Unit: translation.Unit,
		}
	}

	res, err := h.service.TranslateRecipe(c, &api.TranslateRecipeRequest{
		RecipeId:     c.Param(ParamRecipeId),
		TranslatorId: payload.UserId.String(),
		Language:     body.Language,
		Name:         body.Name,
		Description:  body.Description,
		Ingredients:  ingredients,
		Cooking:      body.Cooking,
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}

func (h *Handler) DeleteRecipeTranslation(c *gin.Context) {
	payload, err := request.GetUserPayloadOrResponse(c)
	if err != nil {
		return
	}

	res, err := h.service.DeleteRecipeTranslation(c, &api.DeleteRecipeTranslationRequest{
		RecipeId:    c.Param(ParamRecipeId),
		RequesterId: payload.UserId.String(),
		Language:    c.Param(ParamLanguageCode),
	})
	if err != nil {
		response.FailGrpc(c, err)
		return
	}

	response.Message(c, res.Message)
}
