package logging

import (
	"context"
	"time"

	"github.com/mephistolie/chefbook-backend-common/log"
)

type Events struct{}

func NewEvents() Events {
	return Events{}
}

func (Events) ConfigLoaded(ctx context.Context) {
	log.Log(ctx, log.Event{
		Event:     "config.loaded",
		Message:   "service configuration loaded",
		Component: "config",
	})
}

func (Events) ServiceDependenciesInitializationFailed(ctx context.Context, err error) {
	log.LogFatal(ctx, log.Event{
		Event:     "api_gateway.services.init_failed",
		Message:   "service dependencies initialization failed",
		Component: "app",
		Operation: "init_services",
	}, err)
}

func (Events) AuthMiddlewareInitializationFailed(ctx context.Context, err error) {
	log.LogFatal(ctx, log.Event{
		Event:     "api_gateway.auth_middleware.init_failed",
		Message:   "authentication middleware initialization failed",
		Component: log.ComponentHTTP,
		Operation: "init_auth_middleware",
	}, err)
}

func (Events) HTTPServerFailed(ctx context.Context, err error) {
	log.LogError(ctx, log.Event{
		Event:     "http.server.failed",
		Message:   "HTTP server stopped with an error",
		Component: log.ComponentHTTP,
	}, err)
}

func (Events) UserPayloadReadFailed(ctx context.Context, err error) {
	log.LogError(ctx, log.Event{
		Event:     "http.user_payload.read_failed",
		Message:   "failed to read authenticated user payload from request context",
		Component: log.ComponentHTTP,
		Operation: "read_user_payload",
	}, err)
}

type KeyRefreshFailure struct {
	Attempt     int
	MaxAttempts int
	RetryAfter  time.Duration
}

func (Events) AccessTokenKeyRefreshFailed(ctx context.Context, data KeyRefreshFailure, err error) {
	log.LogWarnError(ctx, log.Event{
		Event:     "auth.access_token_key.refresh_failed",
		Message:   "failed to retrieve access token signing key",
		Component: log.ComponentHTTP,
		Operation: "refresh_access_token_key",
		Payload: map[string]any{
			"attempt":        data.Attempt,
			"max_attempts":   data.MaxAttempts,
			"retry_after_ms": data.RetryAfter.Milliseconds(),
		},
	}, err)
}

type HTTPRequest struct {
	Duration time.Duration
	Method   string
	Path     string
	Status   int
	Failed   bool
}

func (Events) HTTPRequestCompleted(ctx context.Context, data HTTPRequest) {
	event := log.Event{
		Event:      "http.request.completed",
		Message:    "HTTP request completed",
		Component:  log.ComponentHTTP,
		Duration:   data.Duration,
		HTTPMethod: data.Method,
		HTTPPath:   data.Path,
		HTTPStatus: data.Status,
	}
	if data.Failed {
		log.LogWarn(ctx, event)
		return
	}
	log.Log(ctx, event)
}
