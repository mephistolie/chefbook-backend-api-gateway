package response_body

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	api "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
	"time"
)

func ShoppingListUsers(users []*api.ShoppingListUser) []response.ProfileInfo {
	var dtos []response.ProfileInfo
	for _, user := range users {
		dtos = append(dtos, response.ProfileInfo{
			Id:     user.Id,
			Name:   user.Name,
			Avatar: user.Avatar,
		})
	}

	return dtos
}

type GetShoppingListLink struct {
	Link                string    `json:"link"`
	ExpirationTimestamp time.Time `json:"expirationTimestamp"`
}
