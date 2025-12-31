package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoggerMiddleware() gin.HandlerFunc {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		duration := time.Since(start)

		logEvent := log.Info()

		logEvent.
			Str("method", c.Request.Method).
			Str("path", path).
			Int("status", c.Writer.Status()).
			Dur("duration_ms", duration).
			Str("ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent())

		if raw != "" {
			logEvent.Str("query", raw)
		}

		logEvent.Int("size", c.Writer.Size())

		if len(c.Errors) > 0 {
			logEvent = log.Error()
			for _, err := range c.Errors {
				logEvent.Err(err)
			}
		}

		logEvent.Msg("HTTP Request")
	}
}
