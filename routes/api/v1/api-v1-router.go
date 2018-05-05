package v1Router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/anant-sharma/go-boilerplate/routes/api/v1/clock"
)

/*
RegisterRouter Definition
*/
func RegisterRouter(router *gin.RouterGroup) {
	clock.RegisterRouter(router.Group("/clock"))
}
