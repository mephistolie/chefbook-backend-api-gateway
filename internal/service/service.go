package service

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
)

type Services struct {
	Auth         *Auth
	User         *User
	Profile      *Profile
	Category     *Category
	ShoppingList *ShoppingList
}

func NewServices(cfg *config.Config) (*Services, error) {
	authService, err := NewAuth(*cfg.AuthService.Addr)
	if err != nil {
		return nil, err
	}
	userService, err := NewUser(*cfg.UserService.Addr)
	if err != nil {
		return nil, err
	}
	profileService, err := NewProfile(*cfg.ProfileService.Addr)
	if err != nil {
		return nil, err
	}
	categoryService, err := NewCategory(*cfg.CategoryService.Addr)
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
		Profile:      profileService,
		Category:     categoryService,
		ShoppingList: shoppingList,
	}, nil
}

func (s *Services) Stop() error {
	_ = s.Auth.Conn.Close()
	_ = s.User.Conn.Close()
	_ = s.Profile.Conn.Close()
	_ = s.Category.Conn.Close()
	_ = s.ShoppingList.Conn.Close()
	return nil
}
