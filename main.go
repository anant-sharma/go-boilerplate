package main

import (
	"log"

	router "github.com/anant-sharma/go-boilerplate/routes"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	router.InitRouter()

}
