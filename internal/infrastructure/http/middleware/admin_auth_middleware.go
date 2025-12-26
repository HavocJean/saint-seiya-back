package middleware

import (
	"net/http"
	"saint-seiya-back/internal/config"
	"saint-seiya-back/internal/responses"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responses.Error(c, http.StatusUnauthorized, "Authorization header required", "")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower((parts[0])) != "bearer" {
			responses.Error(c, http.StatusUnauthorized, "Token must be in Bearer", "")
			c.Abort()
			return
		}

		token := parts[1]
		expectedToken := config.Cfg.AdminToken

		if token != expectedToken {
			responses.Error(c, http.StatusUnauthorized, "Unathorized", "Invalid admin token")
			c.Abort()
			return
		}

		c.Next()
	}
}
