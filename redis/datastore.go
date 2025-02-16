package redis

import (
	"context"
	"dora-dev-test/data"
)

type DataStore struct {
	// TODO: implement me
}

func (d DataStore) SaveTick(ctx context.Context, tick data.Tick) error {
	//TODO implement me
	panic("implement me")
}

func (d DataStore) GetTicks(ctx context.Context, assetID string, from, to *int64, limit int) ([]data.Tick, error) {
	//TODO implement me
	panic("implement me")
}

func NewDataStore() *DataStore {
	return &DataStore{}
}
