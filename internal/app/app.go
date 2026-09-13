package app

import (
	"context"
	"errors"

	"github.com/mephistolie/chefbook-backend-api-gateway/internal/config"
	eventlog "github.com/mephistolie/chefbook-backend-api-gateway/internal/logging"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	events := eventlog.NewEvents()
	cfg.Print(ctx)

	services, err := service.NewServices(cfg)
	if err != nil {
		events.ServiceDependenciesInitializationFailed(ctx, err)
		return
	}

	authMiddleware, err := auth.NewMiddleware(ctx, services.Auth, *cfg.AuthService.AccessTokenKeyUpdateInterval)
	if err != nil {
		events.AuthMiddlewareInitializationFailed(ctx, err)
		return
	}

	h := handler.NewHandler(services, cfg)
	r := router.NewRouter(h, authMiddleware)

	srv := server.NewServer(*cfg.Port, r.Init(cfg))

	go runServer(ctx, srv)

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

func runServer(ctx context.Context, srv *server.Server) {
	if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
		eventlog.NewEvents().HTTPServerFailed(ctx, err)
	}
}
