package middleware

import (
	"net/http"
	"saint-seiya-back/internal/infrastructure/services"
	"saint-seiya-back/internal/responses"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJwtMiddleware(jwtService *services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responses.Error(c, http.StatusUnauthorized, "Authorization header required", "")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			responses.Error(c, http.StatusUnauthorized, "Token must be in Bearer format", "")
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			responses.Error(c, http.StatusUnauthorized, "Invalid or expired token", err.Error())
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if userID, ok := claims["user_id"].(float64); ok {
			c.Set("user_id", uint(userID))
		}

		c.Next()
	}
}
