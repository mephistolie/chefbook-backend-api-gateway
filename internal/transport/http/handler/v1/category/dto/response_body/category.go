package response_body

import (
	api "github.com/mephistolie/chefbook-backend-category/api/proto/implementation/v1"
)

type Category struct {
	Id    string  `json:"categoryId" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Emoji *string `json:"emoji,omitempty"`
}

func GetCategories(response *api.GetUserCategoriesResponse) []Category {
	categories := make([]Category, len(response.Categories))
	for i, category := range response.Categories {
		categories[i] = Category{
			Id:    category.CategoryId,
			Name:  category.Name,
			Emoji: category.Emoji,
		}
	}
	return categories
}

type AddCategory struct {
	Id string `json:"categoryId" binding:"required"`
}
