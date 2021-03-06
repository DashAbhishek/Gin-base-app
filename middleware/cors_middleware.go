package middleware

import (
	"base-app/utils"

	"github.com/gin-gonic/gin"
)

// standard cors middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		Logger := utils.Logger
		Logger.Info("cors middleware: intercepted")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
