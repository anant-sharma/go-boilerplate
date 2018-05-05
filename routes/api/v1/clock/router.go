package clock

import (
	"github.com/gin-gonic/gin"
)

/*
RegisterRouter Router Group
*/
func RegisterRouter(router *gin.RouterGroup) {
	router.GET("/", GetTime)
}
