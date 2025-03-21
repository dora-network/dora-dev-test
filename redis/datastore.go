package redis

import (
	"context"
	"dora-dev-test/data"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type DataStore struct {
	// TODO: implement me
	client *redis.Client
}

func (d DataStore) SaveTick(ctx context.Context, tick data.Tick) error {
	//TODO implement me
	value, _ := json.Marshal(tick)
	err := d.client.Set(ctx, tick.AssetID, string(value), 0).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func (d DataStore) GetTicks(ctx context.Context, assetID string, from, to *int64, limit int) ([]data.Tick, error) {
	//TODO implement me
	val, err := d.client.Get(ctx, assetID).Result()
	if err != nil {
		panic(err)
	}
	var tick data.Tick
	err = json.Unmarshal([]byte(val), &tick)
	if err != nil {
		return nil, err
	}
	return []data.Tick{tick}, nil
}

func NewDataStore() *DataStore {
	return &DataStore{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}
