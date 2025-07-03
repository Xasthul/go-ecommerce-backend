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

func (h *APIHandler) RegisterHandler(c *gin.Context) {}

func (h *APIHandler) LoginHandler(c *gin.Context) {}
