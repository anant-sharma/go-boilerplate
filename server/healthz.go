package server

import (
	"context"

	"github.com/anant-sharma/go-boilerplate/protos"
	opentracing "github.com/anant-sharma/go-utils/open-tracing"
)

func (*server) GetHealth(ctx context.Context, _ *protos.GetHealthRequest) (*protos.GetHealthResponse, error) {
	span := opentracing.CreateChildSpanFromContext(ctx, "server.GetHealth")
	defer span.Finish()

	return &protos.GetHealthResponse{
		IsHealthy: true,
	}, nil
}
