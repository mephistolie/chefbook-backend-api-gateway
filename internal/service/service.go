package service

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
)

type Services struct {
	Auth         *Auth
	User         *User
	ShoppingList *ShoppingList
}

func NewServices(cfg *config.Config) (*Services, error) {
	authService, err := NewAuth(*cfg.AuthService.Addr)
	userService, err := NewUser(*cfg.UserService.Addr)
	if err != nil {
		return nil, err
	}
	shoppingList, err := NewShoppingList(*cfg.ShoppingListService.Addr)
	if err != nil {
		return nil, err
	}

	return &Services{
		Auth:         authService,
		User:         userService,
		ShoppingList: shoppingList,
	}, nil
}

func (s *Services) Stop() error {
	_ = s.Auth.Conn.Close()
	_ = s.User.Conn.Close()
	_ = s.ShoppingList.Conn.Close()
	return nil
}
