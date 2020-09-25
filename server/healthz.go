package server

import (
	"context"

	"github.com/anant-sharma/go-boilerplate/protos"
)

func (*server) GetHealth(ctx context.Context, _ *protos.GetHealthRequest) (*protos.GetHealthResponse, error) {
	return &protos.GetHealthResponse{
		IsHealthy: true,
	}, nil
}
