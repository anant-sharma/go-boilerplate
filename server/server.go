package server

import (
	"os"
	"path/filepath"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/protos"
	"google.golang.org/grpc"

	serverUtils "github.com/anant-sharma/go-boilerplate-utils/server"
)

type server protos.UnimplementedClockServiceServer

// Start - Method to Start gRPC server
func Start() {
	_, grpcServer, lis := serverUtils.Start(serverUtils.Config(config.C), func(grpcServer *grpc.Server) {
		protos.RegisterClockServiceServer(grpcServer, &server{})
	})

	gwServer := serverUtils.StartHTTPProxy(serverUtils.Config(config.C), getOpenAPIPath(), protos.RegisterClockServiceHandlerFromEndpoint)

	serverUtils.StopWatch(grpcServer, lis, gwServer)
}

func getOpenAPIPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(dir, "/third_party/OpenAPI")
}
