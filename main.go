package main

import (
	"log"
	"strconv"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	/*
		Get Application Configuration
	*/
	config := config.GetConfig()

	/*
		Create Instance of Router
	*/
	r := gin.Default()

	/*
		Register Router
	*/
	router.Register(r.Group("/"))

	/*
		Start Server
	*/
	err := r.Run(":" + strconv.Itoa(config.PORT))
	if err != nil {
		log.Fatal(err)
	}

}
