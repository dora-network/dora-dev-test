package service

import (
	"context"
	api "dora-dev-test/api/v1"
	"dora-dev-test/redis"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	api.UnimplementedDoraDevTestServiceServer

	ds *redis.DataStore
}

func (s Service) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*api.HealthCheckResponse, error) {
	var response api.HealthCheckResponse
	response.LastHeartbeat = timestamppb.Now()
	return &response, nil
}

func (s Service) GetTicks(ctx context.Context, request *api.GetTicksRequest) (*api.GetTicksResponse, error) {
	ticks, err := s.ds.GetTicks(ctx, request.Symbol, nil, nil, 0)
	if err != nil {
		return nil, err
	}
	var response api.GetTicksResponse
	for _, tick := range ticks {
		var apiTick api.Tick
		apiTick.AssetId = tick.AssetID
		apiTick.LastPrice = tick.LastPrice
		response.Ticks = append(response.Ticks, &apiTick)
	}
	return &response, nil
}

func NewService(ds *redis.DataStore) Service {
	return Service{ds: ds}
}
