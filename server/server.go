package server

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/protos"
	"github.com/anant-sharma/go-utils"
	opentracing "github.com/anant-sharma/go-utils/open-tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// Static files
	_ "github.com/anant-sharma/go-boilerplate/statik"
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

	s := grpc.NewServer(
		withServerUnaryInterceptor(),
	)
	protos.RegisterClockServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Starting Server on", hostAddress)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal("Unable to start server: ", err)
			return
		}
	}()

	go func() {
		log.Panicln("Unable to start HTTP Proxy", startHTTPProxy(hostAddress))
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

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// Start Open Tracing Span
	reqID := utils.GenerateShortID()
	span, c := opentracing.StartSpanFromContext(ctx, reqID)
	span.SetTag("X-Request-ID", reqID)
	span.SetTag("Method", info.FullMethod)
	defer span.Finish()

	// Calls the handler
	h, err := handler(c, req)

	return h, err
}
