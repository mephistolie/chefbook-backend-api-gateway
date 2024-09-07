package request_body

type RateRecipe struct {
	Score *int32 `json:"score,omitempty"`
}

type SetRecipeCollections struct {
	Collections *[]string `json:"collections,omitempty"`
}
