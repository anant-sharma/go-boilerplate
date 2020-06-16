package main

import (
	"log"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config.Load()

	quit := make(chan bool, 1)

	go func() {
		server.Start()
		quit <- true
	}()

	<-quit
	log.Println("Exiting Main ...")
}
