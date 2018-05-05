package apiRouter

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/anant-sharma/go-boilerplate/routes/api/v1"
)

/*
RegisterRouter Definition
*/
func RegisterRouter(router *gin.RouterGroup) {
	v1Router.RegisterRouter(router.Group("/v1"))
}
