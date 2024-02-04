package common_body

type IngredientItem struct {
	Id       string   `json:"id"`
	Text     *string  `json:"text,omitempty"`
	Type     string   `json:"type"`
	Amount   *float32 `json:"amount,omitempty"`
	Unit     *string  `json:"unit,omitempty"`
	RecipeId *string  `json:"recipeId,omitempty"`
}
