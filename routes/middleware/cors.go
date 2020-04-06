package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Ref: https://github.com/gin-contrib/cors

// Cors function definition
func Cors() gin.HandlerFunc {
	return cors.Default()
}
