package server

import (
	"context"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/anant-sharma/go-boilerplate/config"
	"github.com/anant-sharma/go-boilerplate/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// Static files
	_ "github.com/anant-sharma/go-boilerplate/third_party/statik"
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

func startHTTPProxy(endpoint string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register Service
	err := proto.RegisterClockServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
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
		Addr: hostAddress,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/swagger") {

				if r.URL.Path == "/swagger" {
					http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
					return
				}

				oaHandler.ServeHTTP(w, r)
				return
			}

			mux.ServeHTTP(w, r)
		}),
	}

	log.Printf("Serving gRPC-Gateway on http://%s", hostAddress)
	log.Printf("Serving OpenAPI Documentation on http://%s/swagger/", hostAddress)

	gwErr := gwServer.ListenAndServe()
	if gwErr != nil {
		return gwErr
	}

	return nil
}

// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")

	statikFS, err := fs.New()
	if err != nil {
		panic("creating OpenAPI filesystem: " + err.Error())
	}

	return http.FileServer(statikFS)
}
