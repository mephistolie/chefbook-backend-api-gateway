package common_body

type Purchase struct {
	Id          string  `json:"purchaseId" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Multiplier  *int32  `json:"multiplier,omitempty"`
	Purchased   bool    `json:"purchased"`
	Amount      *int32  `json:"amount,omitempty"`
	MeasureUnit *string `json:"measureUnit,omitempty"`
	RecipeId    *string `json:"recipeId,omitempty"`
}
