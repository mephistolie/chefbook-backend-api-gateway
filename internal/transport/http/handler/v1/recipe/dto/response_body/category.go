package response_body

import (
	api "github.com/mephistolie/chefbook-backend-recipe/api/proto/implementation/v1"
)

type Category struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Emoji *string `json:"emoji,omitempty"`
}

type CategoryInfo struct {
	Name  string  `json:"name"`
	Emoji *string `json:"emoji,omitempty"`
}

func newCategories(response []*api.RecipeCategory) []Category {
	categories := make([]Category, len(response))
	for i, category := range response {
		categories[i] = newCategory(category)
	}
	return categories
}

func newCategoriesMap(response map[string]*api.RecipeCategoryInfo) map[string]Category {
	categories := make(map[string]Category)
	for id, category := range response {
		categories[id] = newCategoryInfo(category)
	}
	return categories
}

func newCategory(response *api.RecipeCategory) Category {
	return Category{
		Id:    response.Id,
		Name:  response.Name,
		Emoji: response.Emoji,
	}
}

func newCategoryInfo(response *api.RecipeCategoryInfo) Category {
	return Category{
		Name:  response.Name,
		Emoji: response.Emoji,
	}
}
