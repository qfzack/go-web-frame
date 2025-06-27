package middleware

import "github.com/gin-gonic/gin"

// TODO: add tracing middleware
func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Opentelemetry, Jaeger, Zipkin, etc.
	}
}
