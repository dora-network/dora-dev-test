package service

import (
	"context"
	api "dora-dev-test/api/v1"
	"dora-dev-test/datastore"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	api.UnimplementedDoraDevTestServiceServer
	ds datastore.DataStore
}

func (s Service) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*api.HealthCheckResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetTicks(ctx context.Context, request *api.GetTicksRequest) (*api.GetTicksResponse, error) {
	ticks, err := s.ds.GetTicks(ctx, request.GetSymbol(), nil, nil, 1)
	if err != nil || len(ticks) == 0 {
		return nil, err
	}
	// TODO Add ticks when it would possible
	tick := ticks[0]
	resp := &api.GetTicksResponse{
		Ticks: []*api.Tick{{
			AssetId:   tick.AssetID,
			Timestamp: timestamppb.New(tick.Timestamp),
			LastPrice: tick.LastPrice,
			LastSize:  tick.LastSize,
			BestBid:   tick.BestBid,
		}},
	}

	return resp, nil
}

func NewService(ds datastore.DataStore) Service {
	return Service{
		ds: ds,
	}
}
