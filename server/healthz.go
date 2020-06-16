package server

import (
	"context"

	"github.com/anant-sharma/go-boilerplate/proto"
)

func (*server) GetHealth(ctx context.Context, _ *proto.GetHealthRequest) (*proto.GetHealthResponse, error) {
	return &proto.GetHealthResponse{
		IsHealthy: true,
	}, nil
}
