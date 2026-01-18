package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"logging-best-practices/pkg/logger"
)

const (
	// Header del trace-id
	TraceIDHeader string = "X-Trace-ID"
)

// generateTraceID genera un trace-id único
// Formato: 16 bytes en hexadecimal = 32 caracteres
// Ejemplo: "a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4"
func generateTraceID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "fallback-" + time.Now().Format("20060102150405.000000")
	}
	return hex.EncodeToString(bytes)
}

// Logger es el middleware principal de logging
// Se encarga de:
// 1. Generar o propagar el trace-id
// 2. Inyectarlo en el contexto
// 3. Loggear la request de entrada y de salida
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		traceID := c.GetHeader(TraceIDHeader)
		if traceID == "" {
			traceID = generateTraceID()
		}

		ctx := logger.WithTraceID(c.Request.Context(), traceID)
		c.Request = c.Request.WithContext(ctx)
		c.Set("trace_id", traceID)
		c.Header(TraceIDHeader, traceID)

		logger.Info(ctx, "request started",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		)

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		fields := []zap.Field{
			zap.Int("status", statusCode),
			zap.Duration("duration", duration),
			zap.Int("body_size", c.Writer.Size()),
		}

		if len(c.Errors) > 0 {
			fields = append(fields, zap.Strings("errors", c.Errors.Errors()))
		}

		// Usar helper según status code
		if statusCode >= 500 {
			logger.Error(ctx, "request completed", fields...)
		} else if statusCode >= 400 {
			logger.Warn(ctx, "request completed", fields...)
		} else {
			logger.Info(ctx, "request completed", fields...)
		}
	}
}

// Recovery es un middleware de recuperación de panics
// Importante: loggea el panic con el trace-id antes de devolver 500
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx := c.Request.Context()
				logger.Error(ctx, "panic recovered",
					zap.Any("error", err),
					zap.Stack("stacktrace"),
				)
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
