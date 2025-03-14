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
	message := "Test Save!"
	fmt.Println(message)

	return c.ds.SaveTick(ctx, tick)
}

func (c *consumer) Start(parent context.Context) {
	ctx, cancel := context.WithCancel(parent)
	c.cancel = cancel
	message := "Test!"
	fmt.Println(message)
	//client.ConsumeTopics("incoming_prices")
	c.client.AddConsumeTopics("incoming_prices")

	go c.start(ctx)
}

func (c *consumer) start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// TODO implement me
			// Consumer should consume messages from the Kafka topic using the Kafka client
			// and save the tick data to the data store
			fetches := c.client.PollFetches(context.Background())
			if fetches.IsClientClosed() {
				return
			}
			fetches.EachError(func(t string, p int32, err error) {
				//die("fetch err topic %s partition %d: %v", t, p, err)
				//message := "Test!"
				fmt.Println("Got " + t)
			})
			fetches.EachRecord(func(r *kgo.Record) {
				fmt.Print("XXX")
				fmt.Println(r)
				var tick data.Tick
				err := json.Unmarshal(r.Value, &tick)
				if err != nil {
				}
				c.Save(ctx, tick)
			})

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
