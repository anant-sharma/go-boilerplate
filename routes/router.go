package router

import (
	"log"
	"net/http"
	"strconv"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/controller/health"
	"github.com/anant-sharma/go-boilerplate/routes/middleware"
	v1Router "github.com/anant-sharma/go-boilerplate/routes/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter Definition
func InitRouter() {

	config := config.GetConfig()

	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, health.Check(c))
	})

	r.Use(middleware.AuthenticationRequired())
	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1")
	v1Router.InitRouter(v1)

	// Start Server
	err := r.Run(":" + strconv.Itoa(config.PORT))
	if err != nil {
		log.Fatal(err)
	}

}
