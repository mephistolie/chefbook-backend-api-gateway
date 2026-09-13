package log

import (
	"github.com/gin-gonic/gin"
	eventlog "github.com/mephistolie/chefbook-backend-api-gateway/internal/logging"

	"time"
)

const unmatchedRoute = "unmatched"

func Middleware(skipPath []string) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(skipPath); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range skipPath {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		// Process request
		c.Next()
		if _, ok := skip[path]; !ok {
			if routePattern := c.FullPath(); routePattern != "" {
				path = routePattern
			} else {
				path = unmatchedRoute
			}

			eventlog.NewEvents().HTTPRequestCompleted(c.Request.Context(), eventlog.HTTPRequest{
				Duration: time.Since(start),
				Method:   c.Request.Method,
				Path:     path,
				Status:   c.Writer.Status(),
				Failed:   c.Writer.Status() >= 500 || len(c.Errors) > 0,
			})
		}
	}
}
