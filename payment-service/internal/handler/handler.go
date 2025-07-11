package handler

import (
	"github.com/Xasthul/go-ecommerce-backend/payment-service/internal/service"
	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
	s *service.PaymentService
}

func NewApiHandler(s *service.PaymentService) *ApiHandler {
	return &ApiHandler{s: s}
}

func (h *ApiHandler) RegisterRoutes(r *gin.Engine) {

}
