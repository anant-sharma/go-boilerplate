package server

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/protos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func gatewayHandler(mux, oaHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/swagger") {

			if r.URL.Path == "/swagger" {
				http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
				return
			}

			oaHandler.ServeHTTP(w, r)
			return
		}

		mux.ServeHTTP(w, r)
	})
}

func startHTTPProxy(endpoint string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register Service
	err := protos.RegisterClockServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	// Open API Handler
	oaHandler := http.StripPrefix("/swagger/", getOpenAPIHandler())

	hostAddress := strings.Join([]string{
		config.C.Http.Host,
		strconv.Itoa(config.C.Http.Port),
	}, ":")

	log.Println("Starting HTTP Proxy on", hostAddress)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	gwServer := &http.Server{
		Addr:    hostAddress,
		Handler: tracing(logging(gatewayHandler(mux, oaHandler))),
	}

	log.Printf("Serving gRPC-Gateway on http://%s", hostAddress)
	log.Printf("Serving OpenAPI Documentation on http://%s/swagger/", hostAddress)

	gwErr := gwServer.ListenAndServe()
	if gwErr != nil {
		return gwErr
	}

	return nil
}
