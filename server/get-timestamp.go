package server

import (
	"context"
	"time"

	"github.com/anant-sharma/go-boilerplate/controller/clock"
	"github.com/anant-sharma/go-boilerplate/protos"

	newrelictracing "github.com/anant-sharma/go-utils/new-relic/tracing"
)

func (*server) GetTimestamp(ctx context.Context, _ *protos.GetTimestampRequest) (*protos.GetTimestampResponse, error) {
	// span := opentracing.CreateChildSpanFromContext(ctx, "server.GetTimestamp")
	// defer span.Finish()

	segment, ctxWithSegment := newrelictracing.NewSegment(ctx, "server.GetTimestamp")
	defer segment.End()

	return &protos.GetTimestampResponse{
		Timestamp: clock.GetTimeStamp(ctxWithSegment).Timestamp.Format(time.RFC3339),
	}, nil
}

func (*server) GetNewTimestamp(ctx context.Context, _ *protos.GetTimestampRequest) (*protos.GetTimestampResponse, error) {
	// span := opentracing.CreateChildSpanFromContext(ctx, "server.GetTimestamp")
	// defer span.Finish()

	segment, ctxWithSegment := newrelictracing.NewSegment(ctx, "server.GetTimestamp")
	defer segment.End()

	return &protos.GetTimestampResponse{
		Timestamp: clock.GetTimeStamp(ctxWithSegment).Timestamp.Format(time.RFC3339),
	}, nil
}
