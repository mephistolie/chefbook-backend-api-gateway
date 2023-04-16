package auth

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/service"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/request"
	"github.com/mephistolie/chefbook-backend-api-gateway/internal/transport/http/helpers/response"
	v1 "github.com/mephistolie/chefbook-backend-auth/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-common/tokens/access"
	"strings"
	"sync"
	"time"
)

type Middleware struct {
	sync.RWMutex
	authService        service.Auth
	tokenParser        *access.Parser
	keyUpdateTimestamp time.Time
	keyUpdateInterval  time.Duration
}

func NewMiddleware(service service.Auth, keyUpdateInterval time.Duration) (*Middleware, error) {
	res, err := service.GetAccessTokenPublicKey(context.Background(), &v1.GetAccessTokenPublicKeyRequest{})
	if err != nil {
		return nil, err
	}
	parser, err := access.NewParserByKey(res.PublicKey)
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

func (m *Middleware) AuthorizeUser(c *gin.Context) {
	payload, err := m.parseAuthHeader(c)
	if err != nil {
		response.Unauthorized(c, err)
	}
	request.PutUserPayload(c, payload)

}

func (m *Middleware) parseAuthHeader(c *gin.Context) (access.Payload, error) {
	parser := m.fetchPublicKey()

	header := c.GetHeader("Authorization")
	if header == "" {
		return access.Payload{}, errors.New("empty Authorization header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return access.Payload{}, errors.New("invalid Authorization header")
	}

	if len(headerParts[1]) == 0 {
		return access.Payload{}, errors.New("empty access token")
	}

	return parser.Parse(headerParts[1])
}

func (m *Middleware) fetchPublicKey() access.Parser {
	m.Lock()
	if time.Now().UnixMilli()-m.keyUpdateTimestamp.UnixMilli() > m.keyUpdateInterval.Milliseconds() {
		if res, err := m.authService.GetAccessTokenPublicKey(context.Background(), &v1.GetAccessTokenPublicKeyRequest{}); err == nil {
			if parser, err := access.NewParserByKey(res.PublicKey); err == nil {
				m.tokenParser = parser
			}
		}
		m.keyUpdateTimestamp = time.Now()
	}
	parser := *m.tokenParser
	m.Unlock()
	return parser
}
