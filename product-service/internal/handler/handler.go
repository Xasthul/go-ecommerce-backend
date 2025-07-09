package handler

import (
	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/middleware"
	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type APIHandler struct {
	productService  *service.ProductService
	categoryService *service.CategoryService
}

func NewAPIHandler(
	productService *service.ProductService,
	categoryService *service.CategoryService,
) *APIHandler {
	return &APIHandler{
		productService:  productService,
		categoryService: categoryService,
	}
}

func (h *APIHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/products", h.getProducts)
	r.GET("/products/:id", h.getProductById)

	admin := r.Group("/admin", middleware.AdminOnly())
	admin.POST("/products", h.createProduct)
	admin.PUT("/products/:id", h.updateProduct)
	admin.DELETE("/products/:id", h.deleteProduct)

	admin.POST("/categories", h.createCategory)
}

func (h *APIHandler) getProducts(c *gin.Context) {
	products, err := h.productService.GetProducts(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(200, products)
}

type getProductByIdRequest struct {
	ProductId uuid.UUID `json:"productId" binding:"required,uuid"`
}

func (h *APIHandler) getProductById(c *gin.Context) {
	var req getProductByIdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	product, err := h.productService.GetProductById(c.Request.Context(), req.ProductId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch product"})
		return
	}
	c.JSON(200, product)
}

func (h *APIHandler) createProduct(c *gin.Context) {}

func (h *APIHandler) updateProduct(c *gin.Context) {}

func (h *APIHandler) deleteProduct(c *gin.Context) {}

type createCategoryRequest struct {
	Name string `json:"category_name" binding:"required"`
}

func (h *APIHandler) createCategory(c *gin.Context) {
	var req createCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	category, err := h.categoryService.CreateCategory(c.Request.Context(), req.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(201, category)
}
