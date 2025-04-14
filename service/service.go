package service

import (
	"context"

	api "dora-dev-test/api/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	api.UnimplementedDoraDevTestServiceServer
}

func (s Service) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*api.HealthCheckResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s Service) GetTicks(ctx context.Context, request *api.GetTicksRequest) (*api.GetTicksResponse, error) {
	// TODO implement me
	panic("implement me")
}

func NewService() Service {
	return Service{}
}
