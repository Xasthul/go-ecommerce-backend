package handler

import (
	"net/http"

	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/middleware"
	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApiHandler struct {
	productService  *service.ProductService
	categoryService *service.CategoryService
}

func NewApiHandler(
	productService *service.ProductService,
	categoryService *service.CategoryService,
) *ApiHandler {
	return &ApiHandler{
		productService:  productService,
		categoryService: categoryService,
	}
}

func (h *ApiHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/products", h.getProducts)
	r.GET("/products/:id", h.getProductById)

	admin := r.Group("/admin", middleware.AdminOnly())
	admin.POST("/products", h.createProduct)
	admin.PATCH("/products/:id", h.updateProduct)
	admin.DELETE("/products/:id", h.deleteProduct)

	admin.POST("/categories", h.createCategory)
}

func (h *ApiHandler) getProducts(c *gin.Context) {
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

func (h *ApiHandler) getProductById(c *gin.Context) {
	var req getProductByIdURI
	if err := c.ShouldBindUri(&req); err != nil {
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

func (h *ApiHandler) createProduct(c *gin.Context) {
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

type updateProductURI struct {
	ProductId string `uri:"id" binding:"required,uuid"`
}

type updateProductBody struct {
	CategoryID  *int16  `json:"category_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PriceCents  *int32  `json:"price_cents,omitempty"  binding:"omitempty,gt=0"`
	Currency    *string `json:"currency,omitempty"  binding:"omitempty,len=3"`
	Stock       *int32  `json:"stock,omitempty"`
}

func (h *ApiHandler) updateProduct(c *gin.Context) {
	var uri updateProductURI
	var req updateProductBody
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	productId, err := uuid.Parse(uri.ProductId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err = h.productService.UpdateProduct(
		c.Request.Context(),
		productId,
		req.CategoryID,
		req.Name,
		req.Description,
		req.PriceCents,
		req.Currency,
		req.Stock,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update product"})
		return
	}
	c.Status(204)
}

type deleteProductURI struct {
	ProductId string `uri:"id" binding:"required,uuid"`
}

func (h *ApiHandler) deleteProduct(c *gin.Context) {
	var uri deleteProductURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	productId, err := uuid.Parse(uri.ProductId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
		return
	}

	err = h.productService.DeleteProduct(c.Request.Context(), productId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete product"})
		return
	}
	c.Status(200)
}

type createCategoryBody struct {
	Name string `json:"category_name" binding:"required"`
}

func (h *ApiHandler) createCategory(c *gin.Context) {
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
