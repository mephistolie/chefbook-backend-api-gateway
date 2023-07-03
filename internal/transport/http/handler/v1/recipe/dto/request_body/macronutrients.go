package request_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler/v1/recipe/dto/common_body"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

func newMacronutrients(request *common_body.Macronutrients) *api.Macronutrients {
	if request == nil {
		return nil
	}
	return &api.Macronutrients{
		Protein:       request.Protein,
		Fats:          request.Fats,
		Carbohydrates: request.Carbohydrates,
	}
}
