package middleware

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(corsConfig cors.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if slices.Contains(corsConfig.AllowOrigins, origin) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else if origin != "" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Origin not allowed"})
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", fmt.Sprintf("%v", corsConfig.AllowCredentials))
		c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(corsConfig.AllowHeaders, ", "))
		c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(corsConfig.AllowMethods, ", "))

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
