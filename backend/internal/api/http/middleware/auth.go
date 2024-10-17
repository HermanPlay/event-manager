package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/HermanPlay/web-app-backend/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtAuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c, cfg)
		if err != nil {
			// Abort the request with the appropriate error code
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		// Continue down the chain to handler etc
		c.Next()
	}
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func TokenValid(c *gin.Context, cfg *config.Config) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(cfg.App.ApiSecret), nil
	})
	if err != nil {
		return err
	}
	return nil
}
