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

func (h *APIHandler) respondWithAppError(c *gin.Context, err error) {
	if appErr, ok := err.(*service.AppError); ok {
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *APIHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/register", h.registerUser)
	r.POST("/login", h.login)
	r.POST("/refresh", h.refreshTokens)
}

type registerBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (h *APIHandler) registerUser(c *gin.Context) {
	var req registerBody
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

type loginBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *APIHandler) login(c *gin.Context) {
	var req loginBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokens, err := h.s.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		h.respondWithAppError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}

type refreshBody struct {
	RefreshToken string `json:"refresh_token" binding:"required,jwt"`
}

func (h *APIHandler) refreshTokens(c *gin.Context) {
	var req refreshBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokens, err := h.s.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		h.respondWithAppError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}
