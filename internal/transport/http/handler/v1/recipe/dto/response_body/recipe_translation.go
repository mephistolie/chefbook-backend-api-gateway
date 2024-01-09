package response_body

import api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"

type RecipeTranslation struct {
	Author Profile `json:"author"`
}

func newRecipeTranslations(translations map[string]*api.RecipeTranslations) map[string][]RecipeTranslation {
	response := map[string][]RecipeTranslation{}
	for language, languageTranslations := range translations {
		var dto []RecipeTranslation
		for _, translation := range languageTranslations.Translations {
			dto = append(dto, RecipeTranslation{
				Author: Profile{
					Id:     translation.AuthorId,
					Name:   translation.AuthorName,
					Avatar: translation.AuthorAvatar,
				},
			})
		}
		response[language] = dto
	}
	return response
}
