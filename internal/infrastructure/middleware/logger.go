package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// StructuredLogger logs a gin HTTP request.
func StructuredLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		logger := log.WithField("client_id", param.ClientIP).
			WithField("method", param.Method).
			WithField("status_code", param.StatusCode).
			WithField("body_size", param.BodySize).
			WithField("path", param.Path).
			WithField("latency", param.Latency.String())

		if c.Writer.Status() >= http.StatusInternalServerError {
			logger.Error(param.ErrorMessage)
		} else {
			logger.Info(param.ErrorMessage)
		}
	}
}
