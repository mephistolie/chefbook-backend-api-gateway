package docs

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mephistolie/chefbook-backend-api-gateway/contracts"
)

func TestReferenceRoutesAndCompressedAsset(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	NewRouter().Init(r.Group("/"))
	get := func(path, encoding string) *httptest.ResponseRecorder {
		req := httptest.NewRequest("GET", path, nil)
		req.Header.Set("Accept-Encoding", encoding)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w
	}
	for _, path := range []string{"/docs", "/docs/", "/docs/index.html"} {
		w := get(path, "")
		if w.Code != 200 || w.Header().Get("Location") != "" || !strings.Contains(w.Body.String(), "/docs/scalar-1.69.0.js") {
			t.Fatalf("reference page %s: %d", path, w.Code)
		}
	}
	if w := get("/openapi.yaml", ""); w.Code != 200 || !bytes.Equal(w.Body.Bytes(), contracts.OpenAPI) {
		t.Fatal("reference must serve the pinned provider contract")
	}
	compressed := get("/docs/scalar-1.69.0.js", "gzip, deflate")
	if compressed.Code != 200 || compressed.Header().Get("Content-Encoding") != "gzip" {
		t.Fatal("missing gzip response")
	}
	reader, err := gzip.NewReader(compressed.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	decoded, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	for _, encoding := range []string{"", "gzip;q=0", "br, gzip;q=0.0"} {
		plain := get("/docs/scalar-1.69.0.js", encoding)
		if plain.Header().Get("Content-Encoding") != "" || !bytes.Equal(plain.Body.Bytes(), decoded) {
			t.Fatalf("invalid fallback for %q", encoding)
		}
	}
	for _, path := range []string{"/doc", "/doc/", "/doc/index.html"} {
		if w := get(path, ""); w.Code != 404 || w.Header().Get("Location") != "" {
			t.Fatalf("removed route %s: status %d, location %q", path, w.Code, w.Header().Get("Location"))
		}
	}
	if w := get("/docs/missing.js", ""); w.Code != 404 {
		t.Fatal("unknown assets must not serve HTML")
	}
}
