package request_body

import "github.com/google/uuid"

type JoinShoppingList struct {
	Key string `json:"key"`
}

type DeleteUserFromShoppingList struct {
	UserId uuid.UUID `json:"userId"`
}
