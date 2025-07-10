package handler

import (
	"github.com/Xasthul/go-ecommerce-backend/order-service/internal/service"
	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
	s *service.OrderService
}

func NewApiHandler(s *service.OrderService) *ApiHandler {
	return &ApiHandler{s: s}
}

func (h *ApiHandler) RegisterRoutes(r *gin.Engine) {}
