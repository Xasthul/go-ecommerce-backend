package handler

import (
	"github.com/Xasthul/go-ecommerce-backend/auth-service/internal/service"
	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	s *service.AuthService
}

func NewAPIHandler(s *service.AuthService) *APIHandler {
	return &APIHandler{s: s}
}

func (h *APIHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/register", h.registerHandler)
	r.POST("/login", h.loginHandler)
}

func (h *APIHandler) registerHandler(c *gin.Context) {}

func (h *APIHandler) loginHandler(c *gin.Context) {}
