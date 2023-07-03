package common_body

type Macronutrients struct {
	Protein       *int32 `json:"protein,omitempty"`
	Fats          *int32 `json:"fats,omitempty"`
	Carbohydrates *int32 `json:"carbohydrates,omitempty"`
}
