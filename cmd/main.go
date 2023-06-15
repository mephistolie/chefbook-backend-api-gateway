package main

import (
	"flag"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/app"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	_ "github.com/mephistolie/chefbook-backend-common/responses/fail"
	"github.com/peterbourgon/ff/v3"
	"os"
	"time"
)

func main() {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	cfg := config.Config{
		Environment: fs.String("environment", "develop", "service environment"),
		Port:        fs.Int("port", 8080, "service port"),
		LogsPath:    fs.String("logs-path", "", "logs file path"),

		Domains: config.Domains{
			Frontend: fs.String("frontend-domain", "chefbook.io", "Frontend domain"),
			Backend:  fs.String("backend-domain", "api.chefbook.io", "Backend domain"),
		},

		Limiter: config.Limiter{
			RPS:   fs.Int("limiter-rps", 10, "Limiter rates per second"),
			Burst: fs.Int("limiter-burst", 2, "Limiter burst"),
			TTL:   fs.Duration("limiter-ttl", 10*time.Minute, "Limiter entries time to life"),
		},

		AuthService: config.AuthService{
			Addr:                         fs.String("auth-addr", "", "auth service address"),
			AccessTokenKeyUpdateInterval: fs.Duration("access-token-key-ttl", 10*time.Minute, "Access token public key fetch interval"),
		},
		UserService: config.UserService{
			Addr: fs.String("user-addr", "", "user service address"),
		},
		ProfileService: config.ProfileService{
			Addr: fs.String("profile-addr", "", "profile service address"),
		},
		TagService: config.TagService{
			Addr: fs.String("tag-addr", "", "tag service address"),
		},
		CategoryService: config.CategoryService{
			Addr: fs.String("category-addr", "", "category service address"),
		},
		ShoppingListService: config.ShoppingListService{
			Addr: fs.String("shopping-list-addr", "", "shopping list service address"),
		},
	}
	if err := ff.Parse(fs, os.Args[1:], ff.WithEnvVars()); err != nil {
		panic(err)
	}

	err := cfg.Validate()
	if err != nil {
		panic(err)
	}

	app.Run(&cfg)
}
