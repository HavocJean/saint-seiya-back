package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error().
					Interface("Panic", err).
					Str("stack", string(debug.Stack())).
					Str("path", c.Request.URL.Path).
					Str("method", c.Request.Method).
					Str("ip", c.ClientIP()).
					Msg("PANIC RECOVERED - Critical error")

				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Internal server error",
					"error":   "An unexpected error occurred",
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}
