package handler

import (
	"net/http"

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
	r.POST("/refresh", h.refreshHandler)
	r.POST("/logout", h.logoutHandler)
}

type registerReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (h *APIHandler) registerHandler(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.s.RegisterUser(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

// type loginReq struct {
// 	Email    string `json:"email" binding:"required,email"`
// 	Password string `json:"password" binding:"required"`
// }

func (h *APIHandler) loginHandler(c *gin.Context) {}

// type refreshReq struct {
// 	RefreshToken string `json:"refresh_token" binding:"required"`
// }

func (h *APIHandler) refreshHandler(c *gin.Context) {}

func (h *APIHandler) logoutHandler(c *gin.Context) {}
