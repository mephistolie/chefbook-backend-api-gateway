package common_body

import (
	"github.com/google/uuid"
)

type Purchase struct {
	Id          uuid.UUID  `json:"purchaseId" binding:"required"`
	Name        string     `json:"name" binding:"required"`
	Multiplier  *int       `json:"multiplier,omitempty"`
	Purchased   bool       `json:"purchased"`
	Amount      *int       `json:"amount,omitempty"`
	MeasureUnit *string    `json:"measureUnit,omitempty"`
	RecipeId    *uuid.UUID `json:"recipeId,omitempty"`
	RecipeName  *string    `json:"recipeName,omitempty"`
}
