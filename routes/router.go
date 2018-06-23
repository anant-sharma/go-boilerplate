package router

import (
	"github.com/gin-gonic/gin"
	"github.com/anant-sharma/go-boilerplate/routes/api"
	"github.com/anant-sharma/go-boilerplate/routes/auth"
)

/*
Register Router Definition
*/
func Register(router *gin.RouterGroup) {
	apiRouter.RegisterRouter(router.Group("/api"))
	auth.RegisterRouter(router.Group("/auth"))
}
