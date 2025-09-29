package datastore

import (
	"context"
	"dora-dev-test/data"
	"time"
)

type DataStore interface {
	// SaveTick saves the tick data to the data store
	SaveTick(ctx context.Context, tick data.Tick) error

	// GetTicks returns the ticks for the given asset id and time range
	// If from is nil, it will return all ticks from the beginning
	// If to is nil, it will return all ticks until now
	// If both from and to are nil, it will return all ticks
	// If limit is greater than 0, it will return at most that many ticks
	GetTicks(ctx context.Context, assetID string, from, to *int64, offset, limit int) ([]data.Tick, error)

	// SaveCandle saves candle data to the data store
	SaveCandle(ctx context.Context, candle data.Candle) error

	// GetCandles returns the requested candles from the data store
	// If from is nil, it will return candles from the earliest time available in the data store
	// If to is nil, it will return all candles until the most recent candle
	// If both from and to are nil, it will return all candles
	// If limit is greater than 0, it will return at most that many candles
	GetCandles(ctx context.Context, candleID string, from, to *int64, granularity time.Duration, offset, limit int) ([]data.Candle, error)
}
