package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

type RecipeTranslation struct {
	Author response.ProfileInfo `json:"author"`
}

func newRecipeTranslations(translations map[string]*api.RecipeTranslations) map[string][]RecipeTranslation {
	dtos := map[string][]RecipeTranslation{}
	for language, languageTranslations := range translations {
		var dto []RecipeTranslation
		for _, translation := range languageTranslations.Translations {
			dto = append(dto, RecipeTranslation{
				Author: response.ProfileInfo{
					Id:     translation.AuthorId,
					Name:   translation.AuthorName,
					Avatar: translation.AuthorAvatar,
				},
			})
		}
		dtos[language] = dto
	}
	return dtos
}
