package shopping_list

import (
	"fmt"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
)

const (
	ParamShoppingListId = "shopping_list_id"
)

type Handler struct {
	service *service.ShoppingList
	routes  Routes
}

type Routes struct {
	JoinShoppingList string
}

func NewHandler(service *service.ShoppingList, cfg config.Domains) *Handler {
	return &Handler{
		service: service,
		routes: Routes{
			JoinShoppingList: fmt.Sprint("https://", *cfg.Frontend, "/shopping-lists/%s/join?key=%s"),
		},
	}
}
