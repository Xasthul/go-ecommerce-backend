package handler

import (
	"net/http"

	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/service"
	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	s *service.ProductService
}

func NewAPIHandler(s *service.ProductService) *APIHandler {
	return &APIHandler{s: s}
}

func (h *APIHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/products", h.getProducts)
	r.GET("/products/:id", h.getProductById)

	admin := r.Group("/admin", adminOnly())
	admin.POST("/products", h.createProduct)
	admin.PUT("/products/:id", h.updateProduct)
	admin.DELETE("/products/:id", h.deleteProduct)

	admin.POST("/categories", h.createCategory)
}

func adminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("X-User-Role")
		if role != "1" { // 1 = admin
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin only"})
			return
		}
		c.Next()
	}
}

func (h *APIHandler) getProducts(c *gin.Context)    {}
func (h *APIHandler) getProductById(c *gin.Context) {}
func (h *APIHandler) createProduct(c *gin.Context)  {}
func (h *APIHandler) updateProduct(c *gin.Context)  {}
func (h *APIHandler) deleteProduct(c *gin.Context)  {}

func (h *APIHandler) createCategory(c *gin.Context) {}
