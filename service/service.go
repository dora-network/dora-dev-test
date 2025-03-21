package service

import (
	"context"
	api "dora-dev-test/api/v1"
	"dora-dev-test/datastore"
	"time"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	api.UnimplementedDoraDevTestServiceServer
	ds datastore.DataStore
}

func (s Service) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*api.HealthCheckResponse, error) {
	//TODO implement me
	return &api.HealthCheckResponse{
		LastHeartbeat: timestamppb.New(time.Now()),
	}, nil
}

func (s Service) GetTicks(ctx context.Context, request *api.GetTicksRequest) (*api.GetTicksResponse, error) {
	//TODO implement me
	ticks, err := s.ds.GetTicks(ctx, request.Symbol, &request.Start.Seconds, &request.End.Seconds, 10)
	if err != nil {
		return nil, err
	}

	respticks := []*api.Tick{}
	for _, tick := range ticks {
		respticks = append(respticks, &api.Tick{
			AssetId:   tick.AssetID,
			Timestamp: timestamppb.New(tick.Timestamp),
			LastPrice: tick.LastPrice,
			LastSize:  tick.LastSize,
		})
	}
	return &api.GetTicksResponse{Ticks: respticks}, nil
}

func NewService(ds datastore.DataStore) Service {
	return Service{
		ds: ds,
	}
}
