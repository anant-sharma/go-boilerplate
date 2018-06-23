package apiRouter

import (
	"github.com/anant-sharma/go-boilerplate/routes/api/v1"
	"github.com/gin-gonic/gin"
)

/*
RegisterRouter Definition
*/
func RegisterRouter(router *gin.RouterGroup) {
	v1Router.RegisterRouter(router.Group("/v1"))
}
