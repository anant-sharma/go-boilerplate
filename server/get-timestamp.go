package server

import (
	"context"
	"time"

	"github.com/anant-sharma/go-boilerplate/controller/clock"
	"github.com/anant-sharma/go-boilerplate/protos"
	opentracing "github.com/anant-sharma/go-utils/open-tracing"
)

func (*server) GetTimestamp(ctx context.Context, _ *protos.GetTimestampRequest) (*protos.GetTimestampResponse, error) {
	span := opentracing.CreateChildSpanFromContext(ctx, "server.GetTimestamp")
	defer span.Finish()

	return &protos.GetTimestampResponse{
		Timestamp: clock.GetTimeStamp(opentracing.ContextWithSpan(ctx, span)).Timestamp.Format(time.RFC3339),
	}, nil
}
