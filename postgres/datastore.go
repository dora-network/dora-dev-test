package postgres

import (
	"context"
	"dora-dev-test/data"
	"time"
)

type DataStore struct {
	// TODO: implement me
}

func (d DataStore) SaveTick(ctx context.Context, tick data.Tick) error {
	// TODO implement me
	panic("implement me")
}

func (d DataStore) GetTicks(ctx context.Context, assetID string, from, to *int64, offset, limit int) ([]data.Tick, error) {
	// TODO implement me
	panic("implement me")
}

func (d DataStore) SaveCandle(ctx context.Context, candle data.Candle) error {
	// TODO implement me
	panic("implement me")
}

func (d DataStore) GetCandles(ctx context.Context, candleID string, from, to *int64, granularity time.Duration, offset, limit int) ([]data.Candle, error) {
	// TODO implement me
	panic("implement me")
}

func NewDataStore() (DataStore, error) {
	return DataStore{}, nil
}
