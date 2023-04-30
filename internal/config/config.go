package config

import (
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

	Domains     Domains
	Limiter     Limiter
	AuthService AuthService
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

func (c Config) Validate() error {
	if *c.Environment != EnvProd {
		*c.Environment = EnvDev
	}
	return nil
}

func (c Config) Print() {
	log.Infof("API GATEWAY CONFIGURATION\n"+
		"Environment: %v\n"+
		"Port: %v\n"+
		"Logs path: %v\n\n"+
		"Frontend Domain: %v\n"+
		"Backend Domain: %v\n\n"+
		"Limiter RPS: %v\n"+
		"Limiter Burst: %v\n"+
		"Limiter TTL: %v\n\n"+
		"Access Token Key Refresh Interval: %v\n\n"+
		"Auth Service Address: %v\n\n",
		*c.Environment, *c.Port, *c.LogsPath,
		*c.Domains.Frontend, *c.Domains.Backend,
		*c.Limiter.RPS, *c.Limiter.Burst, *c.Limiter.TTL,
		*c.AuthService.AccessTokenKeyUpdateInterval,
		*c.AuthService.Addr,
	)
}
