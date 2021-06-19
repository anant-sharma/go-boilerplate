package main

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/server"

	newrelictracing "github.com/anant-sharma/go-utils/new-relic/tracing"
	opentracing "github.com/anant-sharma/go-utils/open-tracing"
)

func main() {
	config.Load()

	otracingErr := opentracing.Init(context.Background(), "go-boilerplate")
	if otracingErr != nil {
		log.Println("Unable To Initialize Opentracing")
	}

	nrTracingErr := newrelictracing.Init(context.Background(), config.C.NewRelic)
	if nrTracingErr != nil {
		log.Println("Unable To Initialize New Relic Tracing!!")
	}

	quit := make(chan bool, 1)

	go func() {
		server.Start()
		quit <- true
	}()

	<-quit
	log.Println("Exiting Main ...")
}
