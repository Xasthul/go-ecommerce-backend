package handler

import (
	"net/http"

	"github.com/Xasthul/go-ecommerce-backend/order-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApiHandler struct {
	s *service.OrderService
}

func NewApiHandler(s *service.OrderService) *ApiHandler {
	return &ApiHandler{s: s}
}

func (h *ApiHandler) respondWithAppError(c *gin.Context, err error) {
	if appErr, ok := err.(*service.AppError); ok {
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *ApiHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/orders", h.CreateOrder)
}

type createOrderBody struct {
	UserId    string `json:"user_id" binding:"required,uuid"`
	ProductId string `json:"product_id" binding:"required,uuid"`
	Quantity  int    `json:"quantity"         binding:"required,gt=0"`
}

func (h *ApiHandler) CreateOrder(c *gin.Context) {
	var req createOrderBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
		return
	}
	productId, err := uuid.Parse(req.ProductId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
		return
	}

	err = h.s.CreateOrder(c.Request.Context(), userId, productId, req.Quantity)
	if err != nil {
		h.respondWithAppError(c, err)
		return
	}
	c.Status(http.StatusCreated)
}
