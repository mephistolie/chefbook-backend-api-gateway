package response_body

import "time"

type GetShoppingListLink struct {
	Link                string    `json:"link"`
	ExpirationTimestamp time.Time `json:"expirationTimestamp"`
}
