package request_body

type RateRecipe struct {
	Score *int32 `json:"score,omitempty"`
}

type SetRecipeCategories struct {
	Categories *[]string `json:"categories,omitempty"`
}
