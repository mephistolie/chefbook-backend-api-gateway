package common_body

type CookingItem struct {
	Id       string    `json:"id"`
	Text     *string   `json:"text,omitempty"`
	Type     string    `json:"type"`
	Time     *int32    `json:"time,omitempty"`
	Pictures *[]string `json:"pictures,omitempty"`
	RecipeId *string   `json:"recipeId,omitempty"`
}
