package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/anant-sharma/go-boilerplate/routes"
)

func main() {

	r := gin.Default()

	router.Register(r.Group("/"))

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
