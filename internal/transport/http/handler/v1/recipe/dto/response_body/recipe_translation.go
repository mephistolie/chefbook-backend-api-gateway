package response_body

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

func newRecipeTranslations(translations map[string]*api.LanguageTranslations) map[string][]string {
	dtos := map[string][]string{}
	for language, languageTranslations := range translations {
		dtos[language] = languageTranslations.Translators
	}
	return dtos
}
