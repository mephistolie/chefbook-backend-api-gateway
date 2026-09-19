package docs

import (
	"bytes"
	"compress/gzip"
	"embed"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/contracts"
)

//go:embed assets/*
var assets embed.FS

var scalarScript = sync.OnceValue(func() []byte {
	compressed, _ := assets.ReadFile("assets/scalar-1.69.0.js.gz")
	r, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		panic("invalid embedded Scalar bundle")
	}
	defer r.Close()
	data, err := io.ReadAll(r)
	if err != nil {
		panic("invalid embedded Scalar bundle")
	}
	return data
})

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Init(api *gin.RouterGroup) {
	api.GET("/openapi.yaml", func(c *gin.Context) {
		c.Data(200, "application/yaml; charset=utf-8", contracts.OpenAPI)
	})
	serveReference := func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Referrer-Policy", "same-origin")
		switch c.Param("any") {
		case "", "/", "/index.html":
			page, _ := assets.ReadFile("assets/index.html")
			c.Header("Cache-Control", "no-cache")
			c.Data(http.StatusOK, "text/html; charset=utf-8", page)
		case "/init.js":
			config, _ := assets.ReadFile("assets/init.js")
			c.Header("Cache-Control", "no-cache")
			c.Data(http.StatusOK, "text/javascript; charset=utf-8", config)
		case "/scalar-1.69.0.js":
			c.Header("Cache-Control", "public, max-age=31536000, immutable")
			c.Header("Vary", "Accept-Encoding")
			c.Header("Content-Type", "text/javascript; charset=utf-8")
			var script []byte
			if acceptsGzip(c.GetHeader("Accept-Encoding")) {
				script, _ = assets.ReadFile("assets/scalar-1.69.0.js.gz")
				c.Header("Content-Encoding", "gzip")
			} else {
				script = scalarScript()
			}
			http.ServeContent(c.Writer, c.Request, "scalar-1.69.0.js", time.Time{}, bytes.NewReader(script))
		default:
			c.Status(http.StatusNotFound)
		}
	}
	api.GET("/docs", serveReference)
	api.GET("/docs/*any", serveReference)
}

func acceptsGzip(header string) bool {
	for _, value := range strings.Split(header, ",") {
		parts := strings.Split(value, ";")
		if strings.TrimSpace(parts[0]) != "gzip" {
			continue
		}
		for _, parameter := range parts[1:] {
			if name, value, ok := strings.Cut(strings.TrimSpace(parameter), "="); ok && name == "q" && strings.Trim(value, "0.") == "" {
				return false
			}
		}
		return true
	}
	return false
}
