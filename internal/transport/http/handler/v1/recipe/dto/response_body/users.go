package response_body

import "time"

type GetShoppingListLink struct {
	Link      string    `json:"link"`
	ExpiresAt time.Time `json:"expiresAt"`
}
