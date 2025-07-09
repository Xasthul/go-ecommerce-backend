package handler

import (
	"fmt"
	"net/http"

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

type getProductByIdURI struct {
	ProductId string `uri:"id" binding:"required,uuid"`
}

func (h *APIHandler) getProductById(c *gin.Context) {
	var req getProductByIdURI
	if err := c.ShouldBindUri(&req); err != nil {
		fmt.Println("Error binding URI:", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	id, err := uuid.Parse(req.ProductId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
		return
	}

	product, err := h.productService.GetProductById(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch product"})
		return
	}
	c.JSON(200, product)
}

type createProductBody struct {
	CategoryID  int16   `json:"category_id" binding:"required"`
	Name        string  `json:"name"         binding:"required"`
	Description *string `json:"description,omitempty"`
	PriceCents  int32   `json:"price_cents"  binding:"required,gt=0"`
	Currency    *string `json:"currency,omitempty"  binding:"omitempty,len=3"`
	Stock       *int32  `json:"stock,omitempty"`
}

func (h *APIHandler) createProduct(c *gin.Context) {
	var req createProductBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := h.productService.CreateProduct(
		c.Request.Context(),
		req.CategoryID,
		req.Name,
		req.Description,
		req.PriceCents,
		req.Currency,
		req.Stock,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create product"})
		return
	}
	c.Status(201)
}

func (h *APIHandler) updateProduct(c *gin.Context) {}

func (h *APIHandler) deleteProduct(c *gin.Context) {}

type createCategoryBody struct {
	Name string `json:"category_name" binding:"required"`
}

func (h *APIHandler) createCategory(c *gin.Context) {
	var req createCategoryBody
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
