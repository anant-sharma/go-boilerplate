package v1router

import (
	"net/http"

	v1controller "github.com/anant-sharma/go-boilerplate/controller/v1"
	"github.com/gin-gonic/gin"
)

// InitClockRouter - Function to initialize clock router
func InitClockRouter(router *gin.RouterGroup) {
	router.GET("", func(c *gin.Context) {
		data := v1controller.GetTimeStamp(c)
		c.JSON(http.StatusOK, data)
	})
}
