package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/anant-sharma/go-boilerplate/config"
	"gitlab.com/anant-sharma/go-boilerplate/routes"
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
