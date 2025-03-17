package consumer

import (
	"context"
	"dora-dev-test/data"
	"dora-dev-test/datastore"
	"encoding/json"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer interface {
	Start(ctx context.Context)
	Save(ctx context.Context, tick data.Tick) error
	Stop()
}

type consumer struct {
	client *kgo.Client
	ds     datastore.DataStore
	cancel context.CancelFunc
}

func (c *consumer) Save(ctx context.Context, tick data.Tick) error {
	return c.ds.SaveTick(ctx, tick)
}

func (c *consumer) Start(parent context.Context) {
	ctx, cancel := context.WithCancel(parent)
	c.cancel = cancel

	go c.start(ctx)
}

func (c *consumer) start(ctx context.Context) {
	fmt.Println("start")
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fetches := c.client.PollFetches(ctx)
			if errs := fetches.Errors(); len(errs) > 0 {
				// All errors are retried internally when fetching, but non-retriable errors are
				// returned from polls so that users can notice and take action.
				panic(fmt.Sprint(errs))
			}

			// We can iterate through a record iterator...
			iter := fetches.RecordIter()
			for !iter.Done() {
				record := iter.Next()
				tick := data.Tick{}

				err := json.Unmarshal(record.Value, &tick)
				if err != nil {
					fmt.Println("Error in JSON unmarshalling from json marshalled object:", err)
					return
				}
				c.ds.SaveTick(ctx, tick)
			}
		}
	}
}

func (c *consumer) Stop() {
	if c.cancel != nil {
		c.cancel()
	}
}

func NewConsumer(client *kgo.Client, ds datastore.DataStore) Consumer {
	return &consumer{client: client, ds: ds}
}
