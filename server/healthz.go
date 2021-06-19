package server

import (
	"context"

	"github.com/anant-sharma/go-boilerplate/protos"
	newrelictracing "github.com/anant-sharma/go-utils/new-relic/tracing"
)

func (*server) GetHealth(ctx context.Context, _ *protos.GetHealthRequest) (*protos.GetHealthResponse, error) {
	// span := opentracing.CreateChildSpanFromContext(ctx, "server.GetHealth")
	// defer span.Finish()

	segment, _ := newrelictracing.NewSegment(ctx, "server.GetHealth")
	defer segment.End()

	return &protos.GetHealthResponse{
		IsHealthy: true,
	}, nil
}
