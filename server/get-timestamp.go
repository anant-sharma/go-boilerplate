package server

import (
	"context"
	"time"

	"github.com/anant-sharma/go-boilerplate/controller/clock"
	"github.com/anant-sharma/go-boilerplate/protos"
)

func (*server) GetTimestamp(ctx context.Context, _ *protos.GetTimestampRequest) (*protos.GetTimestampResponse, error) {
	return &protos.GetTimestampResponse{
		Timestamp: clock.GetTimeStamp(ctx).Timestamp.Format(time.RFC3339),
	}, nil
}
