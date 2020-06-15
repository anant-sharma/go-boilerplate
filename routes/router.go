package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

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

	// Setup Server
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.PORT),
		Handler: r,
	}

	go func() {
		// Start Server
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error Starting Server: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting Down Server ...")

	// Close Server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Error:", err)
	}
	log.Println("Server Shutdown Complete.")

}
