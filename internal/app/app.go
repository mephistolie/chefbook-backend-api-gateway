package app

import (
	"context"
	"errors"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/server"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/handler"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/middleware/auth"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/router"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/shutdown"
	"net/http"
	"time"
)

func Run(cfg *config.Config) {
	log.InitWithService("api-gateway", *cfg.LogsPath, *cfg.Environment == config.EnvDev)
	cfg.Print()
	ctx := context.Background()

	services, err := service.NewServices(cfg)
	if err != nil {
		log.LogFatal(ctx, log.Event{
			Event:     "api_gateway.services.init_failed",
			Message:   "error during service initialization",
			Component: "app",
		}, err)
	}

	authMiddleware, err := auth.NewMiddleware(ctx, services.Auth, *cfg.AuthService.AccessTokenKeyUpdateInterval)
	if err != nil {
		log.LogFatal(ctx, log.Event{
			Event:     "api_gateway.auth_middleware.init_failed",
			Message:   "error during auth middleware initialization",
			Component: log.ComponentHTTP,
		}, err)
	}

	h := handler.NewHandler(services, cfg)
	r := router.NewRouter(h, authMiddleware)

	srv := server.NewServer(*cfg.Port, r.Init(cfg))

	go runServer(srv)

	wait := shutdown.Graceful(ctx, 5*time.Second, map[string]shutdown.Operation{
		"services": func(ctx context.Context) error {
			return services.Stop()
		},
		"http-server": func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	<-wait
}

func runServer(srv *server.Server) {
	if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
		log.LogError(context.Background(), log.Event{
			Event:     "http.server.failed",
			Message:   "error occurred while running http server",
			Component: log.ComponentHTTP,
		}, err)
	}
}
