package server

import (
	"context"
	"time"

	"github.com/anant-sharma/go-boilerplate/controller/clock"
	"github.com/anant-sharma/go-boilerplate/proto"
)

func (*server) GetTimestamp(ctx context.Context, _ *proto.GetTimestampRequest) (*proto.GetTimestampResponse, error) {
	return &proto.GetTimestampResponse{
		Timestamp: clock.GetTimeStamp(ctx).Timestamp.Format(time.RFC3339),
	}, nil
}
