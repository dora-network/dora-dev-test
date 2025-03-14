package redis

import (
	"context"
	"dora-dev-test/data"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type DataStore struct {
	client redis.Client
}

func (d DataStore) SaveTick(ctx context.Context, tick data.Tick) error {
	key := tick.AssetID
	value, _ := json.Marshal(tick)

	err := d.client.Set(ctx, key, string(value), 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (d DataStore) GetTicks(ctx context.Context, assetID string) (data.Tick, error) {
	val, err := d.client.Get(ctx, assetID).Result()

	var tick data.Tick
	err = json.Unmarshal([]byte(val), &tick)
	if err != nil {
		//handle err
	}

	return tick, err
}

func NewDataStore() *DataStore {
	return &DataStore{}
}
