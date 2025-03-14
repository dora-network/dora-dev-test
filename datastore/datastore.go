package datastore

import (
	"context"
	"dora-dev-test/data"
)

type DataStore interface {
	// SaveTick saves the tick data to the data store
	SaveTick(ctx context.Context, tick data.Tick) error
	// GetTicks returns the ticks for the given asset id and time range
	// If from is nil, it will return all ticks from the beginning
	// If to is nil, it will return all ticks until now
	// If both from and to are nil, it will return all ticks
	// If limit is greater than 0, it will return at most that many ticks
	GetTicks(ctx context.Context, assetID string) (data.Tick, error)
}
