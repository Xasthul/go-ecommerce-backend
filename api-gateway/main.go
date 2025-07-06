package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/Xasthul/go-ecomerce-backend/api-gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadEnv()

	r := gin.Default()

	authServiceURL, _ := url.Parse(cfg.AuthServiceURL)
	r.Any("/auth/*proxyPath", gin.WrapH(newSingleHostReverseProxy(authServiceURL, "/auth")))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":" + cfg.Port)
}

func newSingleHostReverseProxy(target *url.URL, prefix string) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(target)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.URL.Path = strings.TrimPrefix(req.URL.Path, prefix)
		if req.URL.Path == "" {
			req.URL.Path = "/"
		}
	}
	// TODO: Add trace header
	return proxy
}
