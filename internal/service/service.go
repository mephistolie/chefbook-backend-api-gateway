package service

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
)

type Services struct {
	Auth         *Auth
	ShoppingList *ShoppingList
}

func NewServices(cfg *config.Config) (*Services, error) {
	authService, err := NewAuth(*cfg.AuthService.Addr)
	shoppingList, err := NewShoppingList(*cfg.ShoppingListService.Addr)
	if err != nil {
		return nil, err
	}

	return &Services{
		Auth:         authService,
		ShoppingList: shoppingList,
	}, nil
}

func (s *Services) Stop() error {
	_ = s.Auth.Conn.Close()
	_ = s.ShoppingList.Conn.Close()
	return nil
}
