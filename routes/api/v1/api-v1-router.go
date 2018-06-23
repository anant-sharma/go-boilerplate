package v1Router

import (
	"github.com/anant-sharma/go-boilerplate/routes/api/v1/clock"
	"github.com/gin-gonic/gin"
)

/*
RegisterRouter Definition
*/
func RegisterRouter(router *gin.RouterGroup) {
	clock.RegisterRouter(router.Group("/clock"))
}
