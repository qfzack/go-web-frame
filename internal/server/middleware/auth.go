package middleware

import "github.com/gin-gonic/gin"

// TODO: add JWT auth middleware
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: JWT token verify, user permission control, token refresh
	}
}

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: limit by IP, limit by user,
	}
}
