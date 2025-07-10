package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalOnly(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyFromHeader := c.GetHeader("X-API-KEY")
		if apiKeyFromHeader != apiKey {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Wrong api key"})
			return
		}
		c.Next()
	}
}
