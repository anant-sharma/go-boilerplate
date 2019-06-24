package v1router

import (
	v1controller "github.com/anant-sharma/go-boilerplate/controller/v1"
	"github.com/gin-gonic/gin"
)

// Router Interface
type Router struct {
}

// InitRouter Function
func InitRouter(v1 *gin.RouterGroup) {
	c1 := v1controller.NewController()

	{
		clock := v1.Group("/clock")
		{
			clock.GET("", c1.GetTimeStamp)
		}
	}
}
