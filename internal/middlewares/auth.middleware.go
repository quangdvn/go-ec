package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdvn/go-ec/pkg/responses"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			responses.ErrorResponse(c, responses.ErrCodeInvalidToken)
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
