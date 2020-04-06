package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Ref: https://medium.com/@dandua98/gin-authentication-middleware-e659965877b6

// AuthenticationRequired function to check user authentication
func AuthenticationRequired(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		// Check for token in authorization header
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader != "" {
			arr := strings.Split(authHeader, " ")
			if len(arr) > 1 {
				token = arr[1]
			}
		}

		if token == "" {
			token = c.Query("token")
		}

		// Abort if token is also nil
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
