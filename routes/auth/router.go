package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
RegisterRouter Router Group
*/
func RegisterRouter(router *gin.RouterGroup) {
	router.POST("/", Authenticate)
}

/*
Authenticate function to authenticate user
*/
func Authenticate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"auth": true,
	})
}
