package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/anant-sharma/go-boilerplate/routes/api"
	"gitlab.com/anant-sharma/go-boilerplate/routes/auth"
)

/*
Register Router Definition
*/
func Register(router *gin.RouterGroup) {
	apiRouter.RegisterRouter(router.Group("/api"))
	auth.RegisterRouter(router.Group("/auth"))
}
