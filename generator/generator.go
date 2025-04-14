package generator

import (
	"context"
	"math/rand"
	"time"

	"dora-dev-test/data"
)

var assetIDs = []string{
	"BTC-USD",
	"ETH-USD",
	"BCH-USD",
	"ETH-BTC",
}

func GenerateTick(ctx context.Context, tickCh chan<- data.Tick) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
			tickCh <- getTick()
		}
	}
}

func getTick() data.Tick {
	lastPrice := 10000.0
	lastSize := 0.1
	bestBid := 9999.0
	bestAsk := 10001.0

	skew := rand.Float64()
	if rand.Intn(2) == 0 {
		lastPrice *= 1 + (skew / 100)
		bestAsk *= 1 + (skew / 100)
		bestBid *= 1 + (skew / 100)
		lastSize *= 1 + (skew / 100)
	} else {
		lastPrice *= 1 - (skew / 100)
		bestBid *= 1 - (skew / 100)
		bestAsk *= 1 - (skew / 100)
		lastSize *= 1 - (skew / 100)
	}

	return data.Tick{
		AssetID:   assetIDs[rand.Intn(len(assetIDs))],
		Timestamp: time.Now(),
		LastPrice: lastPrice,
		LastSize:  lastSize,
		BestBid:   bestBid,
		BestAsk:   bestAsk,
	}
}
