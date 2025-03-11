package redis

import (
	"fmt"
	"context"
	"dora-dev-test/data"
	"github.com/redis/go-redis/v9"
)

type DataStore struct {
	rdb *redis.Client
}

func (d DataStore) SaveTick(ctx context.Context, tick data.Tick) error {
	err := d.rdb.Set(ctx, tick.AssetID, tick.LastPrice, 0).Err()

    if err != nil {
        panic(err)
    }
    return err
}

func (d DataStore) GetTicks(ctx context.Context, assetID string, from, to *int64, limit int) ([]data.Tick, error) {
	val, err := d.rdb.Get(ctx, assetID).Result()

	if err != nil {
        panic(err)
    }
    fmt.Println("key", val)
}

func NewDataStore() *DataStore {
	return &DataStore{rdb: redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })}
}
