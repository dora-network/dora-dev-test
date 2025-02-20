package redis

import (
	"context"
	"dora-dev-test/data"

	"github.com/go-redis/redis/v8"
)

type DataStore struct {
	client *redis.Client
}

func (d DataStore) SaveTick(ctx context.Context, tick data.Tick) error {
	result := d.client.Set(ctx, tick.AssetID, tick.LastPrice, 0)
	return result.Err()
}

func (d DataStore) GetTicks(ctx context.Context, assetID string, from, to *int64, limit int) ([]data.Tick, error) {
	result := d.client.Get(ctx, assetID)
	err := result.Err()
	if err != nil {
		return nil, err
	}

	price, err := result.Float64()
	if err != nil {
		return nil, err
	}

	var tick data.Tick
	tick.AssetID = assetID
	tick.LastPrice = price

	return []data.Tick{tick}, nil
}

func NewDataStore() *DataStore {
	var options redis.Options
	options.Addr = "localhost:6379"
	client := redis.NewClient(&options)
	return &DataStore{client}
}
