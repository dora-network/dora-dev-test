package consumer

import (
	"context"
	"dora-dev-test/data"
	"dora-dev-test/datastore"
	"encoding/json"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"log"
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
			// TODO implement me
			// Consumer should consume messages from the Kafka topic using the Kafka client
			// and save the tick data to the data store
			for {

				fetches := c.client.PollFetches(context.Background())
				if fetches.IsClientClosed() {
					return
				}

				fetches.EachError(func(t string, p int32, err error) {
					// die("fetch err topic %s partition %d: %v", t, p, err)
					log.Printf("consumer topic %s partition %d fetch error: %v", t, p, err)
				})

				var rs []*kgo.Record
				fetches.EachRecord(func(r *kgo.Record) {
					// rs = append(rs, r)
					tick := data.Tick{}
					err := json.Unmarshal(r.Value, &tick)
					if err != nil {
						// TODO: Need to decide what to do with the error here
						// is data malformed or something else?
						return
					}
					err = c.ds.SaveTick(ctx, tick)
					if err != nil {
						// TODO: Implement some error handling here
						// maybe no deadletter atm as we expect to be sequential
						return
					}
				})

				// Autocommit for now
				if err := c.client.CommitRecords(context.Background(), rs...); err != nil {
					fmt.Printf("commit records failed: %v", err)
					continue
				}
				fmt.Printf("committed %d records individually--this demo does this in a naive way by just hanging on to all records, but you could just hang on to the max offset record per topic/partition!\n", len(rs))
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
