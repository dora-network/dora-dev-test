package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"dora-dev-test/data"
	"dora-dev-test/datastore"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer interface {
	Start(ctx context.Context)
	Save(ctx context.Context, tick data.Tick) error
	Stop() error
}

type consumer struct {
	client *kgo.Client
	ds     datastore.DataStore
	cancel context.CancelFunc
	err    error
}

func (c *consumer) Save(ctx context.Context, tick data.Tick) error {
	return c.ds.SaveTick(ctx, tick)
}

func (c *consumer) Start(parent context.Context) {
	ctx, cancel := context.WithCancel(parent)
	c.cancel = cancel

	go c.start(ctx)
}

func (c *consumer) start(ctx context.Context) (err error) {
	defer func() { // save the error so we can also surface it during stop()
		if err != nil {
			c.err = err
		}
	}()
	c.client.AddConsumeTopics("incoming_prices")
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			slog.Debug("polling records")
			fetches := c.client.PollRecords(ctx, 80)
			var err error
			for _, e := range fetches.Errors() {
				err = errors.Join(err, fmt.Errorf("partition %d: topic %s: %w", e.Partition, e.Topic, e.Err))
			}
			if err != nil {
				slog.Error("got an error", "err", err.Error())
				return err
			}
			for iter := fetches.RecordIter(); !iter.Done(); {
				rec := iter.Next()
				var tick data.Tick
				if err := json.Unmarshal(rec.Value, &tick); err != nil {
					return fmt.Errorf("unmarshal %q to %T: %w", rec.Value, tick, err)
				}
				slog.Debug("got tick", "tick", tick)

				if err := c.Save(ctx, tick); err != nil {
					return err
				}

			}
		}
	}
}

func (c *consumer) Stop() error {
	if c.cancel != nil {
		c.cancel()
	}
	return c.err
}

func NewConsumer(client *kgo.Client, ds datastore.DataStore) Consumer {
	return &consumer{client: client, ds: ds}
}
