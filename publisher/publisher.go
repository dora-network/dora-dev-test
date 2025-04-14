package publisher

import (
	"context"
	"encoding/json"

	"dora-dev-test/data"

	"github.com/twmb/franz-go/pkg/kgo"
)

type TickPublisher interface {
	// Start the publisher
	Start(ctx context.Context, tickCh chan data.Tick, topic string)
	// PublishTick publishes the tick data to the Kafka topic
	PublishTick(ctx context.Context, tick data.Tick, topic string) error
	// Stop the publisher
	Stop()
}

func NewTickPublisher(client *kgo.Client, logger kgo.Logger) TickPublisher {
	return &tickPublisher{client: client, logger: logger}
}

type tickPublisher struct {
	client *kgo.Client
	cancel context.CancelFunc
	logger kgo.Logger
}

func (t *tickPublisher) Start(parent context.Context, tickCh chan data.Tick, topic string) {
	if t.cancel != nil {
		return
	}

	ctx, cancel := context.WithCancel(parent)
	t.cancel = cancel

	go t.start(ctx, tickCh, topic)
}

func (t *tickPublisher) start(ctx context.Context, tickCh chan data.Tick, topic string) {
	for {
		select {
		case <-ctx.Done():
			return
		case tick := <-tickCh:
			err := t.PublishTick(ctx, tick, topic)
			if err != nil {
				t.logger.Log(kgo.LogLevelError, "failed to publish tick", err)
			}
		}
	}
}

func (t *tickPublisher) Stop() {
	if t.cancel != nil {
		t.cancel()
	}
}

func (t *tickPublisher) PublishTick(ctx context.Context, tick data.Tick, topic string) error {
	bs, err := json.Marshal(tick)
	if err != nil {
		return err
	}
	err = t.client.ProduceSync(ctx, &kgo.Record{
		Topic: topic,
		Value: bs,
	}).FirstErr()
	if err != nil {
		t.logger.Log(kgo.LogLevelError, "failed to produce tick", "error", err)
		return err
	}
	t.logger.Log(kgo.LogLevelInfo, "tick published", "tick", tick)
	return nil
}
