package redis

import (
	"context"
	"dora-dev-test/data"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type DataStore struct {
	// TODO: implement me
	redisdb *redis.Client
}

func (d DataStore) SaveTick(ctx context.Context, tick data.Tick) error {
	//TODO implement me
	tickJSON, err := json.Marshal(tick)
	if err != nil {
		panic(err)
	}
	key := fmt.Sprintf("tick:%d", tick.AssetID)
	err = d.redisdb.Set(ctx, key, tickJSON, 0).Err()
	if err != nil {
		panic(err)
	}
	return err
}

func (d DataStore) GetTicks(ctx context.Context, assetID string, from, to *int64, limit int) ([]data.Tick, error) {
	val, err := d.redisdb.Get(ctx, assetID).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	var tick data.Tick
	err = json.Unmarshal([]byte(val), &tick)
	if err != nil {
	}
	tickArray := []data.Tick{tick}
	return tickArray, err
}

func NewDataStore() *DataStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &DataStore{
		redisdb: rdb,
	}
}
