package consumer

import (
	"context"
	"dora-dev-test/data"
	"dora-dev-test/datastore"
)

type Consumer interface {
	Start(ctx context.Context, tickCh <-chan data.Tick)
	Save(ctx context.Context, tick data.Tick) error
}

type consumer struct {
	ds datastore.DataStore
}

func (c *consumer) Start(ctx context.Context, tickCh <-chan data.Tick) {
	for {
		select {
		case <-ctx.Done():
			return
		case tick := <-tickCh:
			c.Save(ctx, tick)

			// Add additional cases if necessary
			// TODO: New candles should be saved on the minute every minute
		}
	}
}

func (c *consumer) Save(ctx context.Context, tick data.Tick) error {
	// TODO: Implement the logic to save the tick data to the datastore
	return nil
}

func NewConsumer(ds datastore.DataStore) Consumer {
	return &consumer{
		ds: ds,
	}
}
