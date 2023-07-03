package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

func newMacronutrients(response *api.Macronutrients) *common_body.Macronutrients {
	if response == nil || (response.Protein == nil && response.Fats == nil && response.Carbohydrates == nil) {
		return nil
	}
	return &common_body.Macronutrients{
		Protein:       response.Protein,
		Fats:          response.Fats,
		Carbohydrates: response.Carbohydrates,
	}
}
