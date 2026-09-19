package auth

import (
	"context"
	"errors"
	"github.com/mephistolie/chefbook-backend-common/responses/fail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/gin-gonic/gin"
	eventlog "github.com/mephistolie/chefbook-backend-api-gateway/internal/logging"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	auth "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/tokens/access"

	"strings"
	"sync"
	"time"
)

type Middleware struct {
	sync.RWMutex
	authService        *service.Auth
	tokenParser        *access.Parser
	keyUpdateTimestamp time.Time
	keyUpdateInterval  time.Duration
}

func NewMiddleware(ctx context.Context, service *service.Auth, keyUpdateInterval time.Duration) (*Middleware, error) {
	var res *auth.GetAccessTokenPublicKeyResponse = nil
	var err error

	for i := 0; i < 6; i++ {
		if res, err = service.GetAccessTokenPublicKey(ctx, &auth.GetAccessTokenPublicKeyRequest{}); err == nil {
			break
		} else if i+1 < 6 {
			eventlog.NewEvents().AccessTokenKeyRefreshFailed(ctx, eventlog.KeyRefreshFailure{
				Attempt:     i + 1,
				MaxAttempts: 6,
				RetryAfter:  10 * time.Second,
			}, err)
			if err := waitForRetry(ctx, 10*time.Second); err != nil {
				return nil, err
			}
		}
	}
	if err != nil {
		return nil, err
	}

	parser, err := access.NewParserByRawKey(res.PublicKey)
	if err != nil {
		return nil, err
	}

	return &Middleware{
		authService:        service,
		tokenParser:        parser,
		keyUpdateTimestamp: time.Now(),
		keyUpdateInterval:  keyUpdateInterval,
	}, nil
}

func waitForRetry(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func (m *Middleware) AuthorizeUser(c *gin.Context) {
	m.authorizeUser(c, false)
}

func (m *Middleware) AuthorizeDeletedUser(c *gin.Context) {
	m.authorizeUser(c, true)
}

func (m *Middleware) authorizeUser(c *gin.Context, allowDeleted bool) {
	payload, err := m.parseAuthHeader(c.Request.Context(), c)
	if err != nil {
		response.Unauthorized(c, err)
		return
	}
	if payload.SessionID <= 0 {
		response.Unauthorized(c, errors.New("session-bound access token required"))
		return
	}
	// Account state must be current even when this JWT predates deletion or blocking.
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	if _, err := m.authService.Authentication.ValidateSession(ctx, &auth.AuthPrincipal{AccountId: payload.UserId.String(), SessionId: payload.SessionID}); err != nil {
		if status.Code(err) == codes.Unauthenticated || status.Code(err) == codes.NotFound || fail.ParseGrpc(err).Code == 401 {
			response.Unauthorized(c, errors.New("invalid or revoked session"))
		} else {
			response.FailGrpc(c, err)
		}
		return
	}
	info, err := m.authService.GetAuthInfo(ctx, &auth.GetAuthInfoRequest{Id: payload.UserId.String()})
	if err != nil {
		if status.Code(err) == codes.NotFound || fail.ParseGrpc(err).ErrorType == fail.TypeNotFound {
			response.Unauthorized(c, errors.New("invalid access token"))
		} else {
			response.FailGrpc(c, err)
		}
		return
	}
	revoking := c.Request.Method == http.MethodDelete && (c.FullPath() == "/v1/sessions" || strings.HasPrefix(c.FullPath(), "/v1/sessions/"))
	if info.IsBlocked && !revoking {
		response.Fail(c, fail.Response{Code: 403, ErrorType: "profile_blocked", Message: "account blocked"})
		return
	}
	if !allowDeleted && info.DeletionTimestamp != nil {
		response.Fail(c, response.ProfileDeleting)
		return
	}
	payload.Deleted = info.DeletionTimestamp != nil
	request.PutUserPayload(c, payload)
}

func (m *Middleware) parseAuthHeader(ctx context.Context, c *gin.Context) (access.Payload, error) {
	parser := m.fetchPublicKey(ctx)

	header := c.GetHeader("Authorization")
	if header == "" {
		return access.Payload{}, errors.New("empty Authorization header")
	}

	headerParts := strings.Fields(header)
	if len(headerParts) != 2 || !strings.EqualFold(headerParts[0], "Bearer") {
		return access.Payload{}, errors.New("invalid Authorization header")
	}

	if len(headerParts[1]) == 0 {
		return access.Payload{}, errors.New("empty access token")
	}

	return parser.Parse(headerParts[1])
}

func (m *Middleware) fetchPublicKey(ctx context.Context) access.Parser {
	m.Lock()
	if time.Now().UnixMilli()-m.keyUpdateTimestamp.UnixMilli() > m.keyUpdateInterval.Milliseconds() {
		if res, err := m.authService.GetAccessTokenPublicKey(ctx, &auth.GetAccessTokenPublicKeyRequest{}); err == nil {
			if parser, err := access.NewParserByRawKey(res.PublicKey); err == nil {
				m.tokenParser = parser
			}
		}
		m.keyUpdateTimestamp = time.Now()
	}
	parser := *m.tokenParser
	m.Unlock()
	return parser
}
