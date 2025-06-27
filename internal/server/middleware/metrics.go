package middleware

import "github.com/gin-gonic/gin"

// TODO: add monitoring
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: add monitor metrics
	}
}
