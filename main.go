package main

import (
	"context"
	"log"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/server"

	opentracing "github.com/anant-sharma/go-utils/open-tracing"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config.Load()

	otracingErr := opentracing.Init(context.Background(), "go-boilerplate")
	if otracingErr != nil {
		log.Println("Unable To Initialize Opentracing")
	}

	quit := make(chan bool, 1)

	go func() {
		server.Start()
		quit <- true
	}()

	<-quit
	log.Println("Exiting Main ...")
}
