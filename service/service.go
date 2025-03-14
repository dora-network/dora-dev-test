package service

import (
	"context"
	api "dora-dev-test/api/v1"
	datastore "dora-dev-test/datastore"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	api.UnimplementedDoraDevTestServiceServer
	db datastore.DataStore
}

func (s Service) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*api.HealthCheckResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetTicks(ctx context.Context, request *api.GetTicksRequest) (*api.GetTicksResponse, error) {
	tick, err := s.db.GetTicks(ctx, request.Symbol)
	if err != nil {
		//handle err
	}

	ticks := []*api.Tick{
		&api.Tick{
			AssetId: tick.AssetID,
		},
	}

	apiResponse := api.GetTicksResponse{
		Ticks: ticks,
	}

	return &apiResponse, err
}

func NewService(db datastore.DataStore) Service {
	return Service{
		db: db,
	}
}
