package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* Expected: Bearer <token> */
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"error": "No token"})
			c.Abort()
			return
		}
		/* Extract Token and removes "Bearer " part */
		tokenString := strings.Split(authHeader, " ")[1]
		/* Verifies Signature, expiration and validity */
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("mysecretkey"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalide token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
