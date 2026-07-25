package config

import (
	"context"
	"github.com/mephistolie/chefbook-backend-common/log"
	"time"
)

const (
	EnvDev  = "develop"
	EnvProd = "production"
)

type Config struct {
	Environment *string
	Port        *int
	LogsPath    *string

	Domains             Domains
	Limiter             Limiter
	AuthService         AuthService
	UserService         UserService
	SubscriptionService SubscriptionService
	ProfileService      ProfileService
	TagService          TagService
	RecipeService       RecipeService
	EncryptionService   EncryptionService
	ShoppingListService ShoppingListService
}

type Domains struct {
	Frontend *string
	Backend  *string
}

type Limiter struct {
	RPS   *int
	Burst *int
	TTL   *time.Duration
}

type AuthService struct {
	Addr                         *string
	AccessTokenKeyUpdateInterval *time.Duration
}

type UserService struct {
	Addr *string
}

type SubscriptionService struct {
	Addr *string
}

type ProfileService struct {
	Addr *string
}

type TagService struct {
	Addr *string
}

type RecipeService struct {
	Addr *string
}

type EncryptionService struct {
	Addr *string
}

type ShoppingListService struct {
	Addr *string
}

func (c Config) Validate() error {
	if *c.Environment != EnvProd {
		*c.Environment = EnvDev
	}
	return nil
}

func (c Config) Print() {
	log.Log(context.Background(), log.Event{
		Event:     "config.loaded",
		Message:   "service configuration loaded",
		Component: "config",
	})
}
