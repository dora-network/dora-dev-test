package redis

import (
	"context"
	"dora-dev-test/data"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type DataStore struct {
	rdb *redis.Client
	// TODO: implement me
}

func (d DataStore) SaveTick(ctx context.Context, tick data.Tick) error {
	result, err := json.Marshal(tick)
	if err != nil {
		panic(err)
	}
	err = d.rdb.Set(ctx, tick.AssetID, result, 0).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func (d DataStore) GetTicks(ctx context.Context, assetID string, from, to *int64, limit int) ([]data.Tick, error) {
	val, err := d.rdb.Get(ctx, assetID).Result()
	if err != nil {
		panic(err)
	}

	tick := data.Tick{}
	err = json.Unmarshal([]byte(val), &tick)

	return []data.Tick{tick}, err
}

func NewDataStore() *DataStore {
	return &DataStore{redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})}
}
