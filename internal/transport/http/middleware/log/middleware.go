package log

import (
	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-common/log"
	"time"
)

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
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()
		if _, ok := skip[path]; !ok {
			if raw != "" {
				path = path + "?" + raw
			}

			event := log.Event{
				Event:      "http.request.completed",
				Message:    "http request completed",
				Component:  log.ComponentHTTP,
				Duration:   time.Since(start),
				HTTPMethod: c.Request.Method,
				HTTPPath:   path,
				HTTPStatus: c.Writer.Status(),
			}
			if len(c.Errors) > 0 {
				log.LogWarn(c.Request.Context(), event)
			} else {
				log.Log(c.Request.Context(), event)
			}
		}
	}
}
