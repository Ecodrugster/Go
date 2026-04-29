package main

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"

)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
	if !strings.HasPrefix(auth, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}
		tokenStr := strings.Split(auth, " ")[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(float64)

		c.Set("user_id", int(userID))
		c.Next()
	}
	}
}
