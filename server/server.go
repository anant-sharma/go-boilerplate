package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

// Start - Method to Start gRPC server
func Start() {
	hostAddress := strings.Join([]string{
		config.C.Grpc.Host,
		strconv.Itoa(config.C.Grpc.Port),
	}, ":")

	lis, err := net.Listen("tcp", hostAddress)
	if err != nil {
		log.Fatal("Unable to listen on", hostAddress)
		return
	}

	s := grpc.NewServer()
	proto.RegisterClockServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Starting Server on", hostAddress)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal("Unable to start server: ", err)
			return
		}
	}()

	log.Println("Started Server on", hostAddress)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Stop Server
	log.Println("Shutting Down Server ...")
	s.Stop()

	// Close Listener
	log.Println("Closing Listener ...")
	lis.Close()

	log.Println("Server Shutdown Complete.")
}
