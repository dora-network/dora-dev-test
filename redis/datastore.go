package redis

import (
	"context"
	"dora-dev-test/data"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type DataStore struct {
	client *redis.Client
}

func (d DataStore) SaveTick(ctx context.Context, tick data.Tick) error {
	res := d.client.HSet(ctx, fmt.Sprintf("asset:%s", tick.AssetID), tick)
	return res.Err()
}

func (d DataStore) GetTicks(ctx context.Context, assetID string, from, to *int64, limit int) ([]data.Tick, error) {
	// TODO Get all ticks in future, but at the moment just get tick and don't break interface
	res, err := d.GetTick(ctx, assetID, from, to, limit)
	if err != nil {
		return nil, err
	}
	return []data.Tick{res}, nil
}

func (d DataStore) GetTick(ctx context.Context, assetID string, from, to *int64, limit int) (data.Tick, error) {
	res := d.client.HGetAll(ctx, fmt.Sprintf("asset:%s", assetID))

	if res.Err() != nil {
		return data.Tick{}, fmt.Errorf("err fetch tick: %w", res.Err())
	}

	t := data.Tick{}
	err := res.Scan(&t)
	if err != nil {
		return data.Tick{}, fmt.Errorf("err marshalling tick: %w", err)
	}

	return t, nil
}

func NewDataStore() *DataStore {
	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	return &DataStore{
		client: r,
	}
}
