package service

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
)

type Services struct {
	Auth         *Auth
	User         *User
	Subscription *Subscription
	Profile      *Profile
	Tag          *Tag
	Recipe       *Recipe
	Encryption   *Encryption
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
	subscriptionService, err := NewSubscription(*cfg.SubscriptionService.Addr)
	if err != nil {
		return nil, err
	}
	profileService, err := NewProfile(*cfg.ProfileService.Addr)
	if err != nil {
		return nil, err
	}
	tagService, err := NewTag(*cfg.TagService.Addr)
	if err != nil {
		return nil, err
	}
	recipeService, err := NewRecipe(*cfg.RecipeService.Addr)
	if err != nil {
		return nil, err
	}
	encryptionService, err := NewEncryption(*cfg.EncryptionService.Addr)
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
		Subscription: subscriptionService,
		Profile:      profileService,
		Tag:          tagService,
		Recipe:       recipeService,
		Encryption:   encryptionService,
		ShoppingList: shoppingList,
	}, nil
}

func (s *Services) Stop() error {
	_ = s.Auth.Conn.Close()
	_ = s.User.Conn.Close()
	_ = s.Subscription.Conn.Close()
	_ = s.Profile.Conn.Close()
	_ = s.Tag.Conn.Close()
	_ = s.Recipe.Conn.Close()
	_ = s.Encryption.Conn.Close()
	_ = s.ShoppingList.Conn.Close()
	return nil
}
