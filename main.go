package main

import (
	"log"
	"strconv"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/controller/auth"
	"github.com/anant-sharma/go-boilerplate/controller/v1"
	"github.com/gin-gonic/gin"
)

func main() {

	config := config.GetConfig()

	r := gin.Default()

	/*
		Auth Routing
	*/
	auth := r.Group("/auth")
	{
		auth.POST("", authcontroller.Authenticate)
	}

	c1 := v1controller.NewController()
	v1 := r.Group("/api/v1")
	{
		clock := v1.Group("/clock")
		{
			clock.GET("", c1.GetTimeStamp)
		}
	}

	/*
		Start Server
	*/
	err := r.Run(":" + strconv.Itoa(config.PORT))
	if err != nil {
		log.Fatal(err)
	}

}
