package consumer

import (
	"fmt"
	"context"
	"encoding/json"
	"dora-dev-test/data"
	"dora-dev-test/datastore"
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
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fetches := c.client.PollFetches(ctx)
            fetches.EachRecord(func(record *kgo.Record) {
				fmt.Println("Seen record")
				var tick = data.Tick{}
				err := json.Unmarshal(record.Value, &tick)
				if err != nil {
					fmt.Println("error")
				}

				c.Save(ctx, tick)
			})
			// TODO implement me
			// Consumer should consume messages from the Kafka topic using the Kafka client
			// and save the tick data to the data store
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
