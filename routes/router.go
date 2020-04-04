package router

import (
	"log"
	"net/http"
	"strconv"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/controller/health"
	v1Router "github.com/anant-sharma/go-boilerplate/routes/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter Definition
func InitRouter() {

	config := config.GetConfig()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1Router.InitRouter(v1)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, health.Check(c))
	})

	/*
		Start Server
	*/
	err := r.Run(":" + strconv.Itoa(config.PORT))
	if err != nil {
		log.Fatal(err)
	}

}
