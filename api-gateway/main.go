package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/Xasthul/go-ecomerce-backend/api-gateway/internal/config"
	"github.com/Xasthul/go-ecomerce-backend/api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadEnv()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	authService := createReverseProxy(cfg.AuthServiceURL)
	r.POST("/auth/login", gin.WrapH(stripPrefixAndProxy(authService, "/auth")))
	r.POST("/auth/register", gin.WrapH(stripPrefixAndProxy(authService, "/auth")))
	r.POST("/auth/refresh", gin.WrapH(stripPrefixAndProxy(authService, "/auth")))

	productService := createReverseProxy(cfg.ProductServiceURL)
	r.GET("/products", gin.WrapH(productService))
	r.GET("/products/*proxyPath", gin.WrapH(productService))

	loggedIn := r.Group("/admin", middleware.AuthMiddleware(cfg.JwtSecret))
	loggedIn.POST("/products", gin.WrapH(productService))
	loggedIn.PATCH("/products/*proxyPath", gin.WrapH(productService))
	loggedIn.DELETE("/products/*proxyPath", gin.WrapH(productService))
	loggedIn.POST("/categories", gin.WrapH(productService))

	r.Run(":" + cfg.Port)
}

func createReverseProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}

func stripPrefixAndProxy(proxy *httputil.ReverseProxy, prefix string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
		proxy.ServeHTTP(w, r)
	})
}
