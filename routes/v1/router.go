package v1router

import (
	"github.com/gin-gonic/gin"
)

// InitRouter Function
func InitRouter(v1 *gin.RouterGroup) {

	InitClockRouter(v1.Group("/clock"))

}
