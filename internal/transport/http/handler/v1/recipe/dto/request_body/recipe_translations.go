package request_body

type TranslateRecipe struct {
	Language    string                           `json:"language" binding:"required"`
	Name        string                           `json:"name" binding:"required"`
	Description *string                          `json:"description"`
	Ingredients map[string]IngredientTranslation `json:"ingredients" binding:"required"`
	Cooking     map[string]string                `json:"cooking" binding:"required"`
}

type IngredientTranslation struct {
	Text string  `json:"text" binding:"required"`
	Unit *string `json:"unit,omitempty"`
}
