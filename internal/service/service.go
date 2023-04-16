package service

import (
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
)

type Services struct {
	Auth Auth
}

func NewServices(cfg *config.Config) (*Services, error) {
	authService, err := NewAuth(*cfg.AuthService.Addr)
	if err != nil {
		return nil, err
	}

	return &Services{
		Auth: *authService,
	}, nil
}

func (s *Services) Stop() error {
	return s.Auth.Conn.Close()
}
