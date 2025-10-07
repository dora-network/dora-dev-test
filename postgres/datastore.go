package postgres

import (
	"context"
	"fmt"
	"time"

	"dora-dev-test/data"

	"github.com/jmoiron/sqlx"
)

type DataStore struct {
	db *sqlx.DB
}

func (d DataStore) SaveTick(ctx context.Context, tick data.Tick) error {
	_, err := d.db.ExecContext(
		ctx,
		"INSERT INTO ticks (asset_id, last_price, last_size, best_bid, best_ask,created_at ) VALUES ($1, $2, $3, $4, $5, $6)",
		tick.AssetID,
		tick.LastPrice,
		tick.LastSize,
		tick.BestBid,
		tick.BestAsk,
		tick.Timestamp,
	)
	if err != nil {
		return fmt.Errorf("save tick: %w", err)
	}

	return nil
}

func (d DataStore) GetTicks(ctx context.Context, assetID string, from, to *int64, offset, limit int) ([]data.Tick, error) {
	query := `
select asset_id, last_price, last_size, best_bid, best_ask, created_at
	from ticks
where asset_id = $1 AND created_at between $2 and $3
order by created_at
limit $4
offset $5
`
	var timeFrom, timeTo time.Time
	if from != nil {
		timeFrom = time.Unix(*from, 0)
	}
	if to != nil {
		timeTo = time.Unix(*to, 0)
	} else {
		timeTo = time.Now().Add(time.Hour)
	}

	rows, err := d.db.QueryContext(ctx, query, assetID, timeFrom, timeTo, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("get ticks: %w", err)
	}
	defer rows.Close()

	var ticks []data.Tick
	for rows.Next() {
		t := data.Tick{}
		if err = rows.Scan(&t.AssetID, &t.LastPrice, &t.LastSize, &t.BestBid, &t.BestAsk, &t.Timestamp); err != nil {
			return nil, fmt.Errorf("get ticks: %w", err)
		}
		ticks = append(ticks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get ticks: %w", err)
	}

	return ticks, nil
}

func (d DataStore) SaveCandle(ctx context.Context, candle data.Candle) error {
	_, err := d.db.ExecContext(
		ctx,
		"INSERT INTO candles (asset_id, start_ts, open, high, low, close, volume) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		candle.AssetID,
		candle.StartTimestamp,
		candle.Open,   // first
		candle.High,   // max
		candle.Low,    // min
		candle.Close,  // last
		candle.Volume, // sum
	)
	if err != nil {
		return fmt.Errorf("save candle: %w", err)
	}

	return nil
}

// 1m, 5m, 10m, 30m, 1h, 24h
func (d DataStore) GetCandles(ctx context.Context, candleID string, from, to *int64, granularity time.Duration, offset, limit int) ([]data.Candle, error) {
	query := `
select asset_id, 
	from ticks
where asset_id = $1 AND created_at between $2 and $3
order by created_at
limit $4
offset $5
`
	var timeFrom, timeTo time.Time
	if from != nil {
		timeFrom = time.Unix(*from, 0)
	}
	if to != nil {
		timeTo = time.Unix(*to, 0)
	} else {
		timeTo = time.Now().Add(time.Hour)
	}

	rows, err := d.db.QueryContext(ctx, query, assetID, timeFrom, timeTo, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("get ticks: %w", err)
	}
	defer rows.Close()

	var ticks []data.Tick
	for rows.Next() {
		t := data.Tick{}
		if err = rows.Scan(&t.AssetID, &t.LastPrice, &t.LastSize, &t.BestBid, &t.BestAsk, &t.Timestamp); err != nil {
			return nil, fmt.Errorf("get ticks: %w", err)
		}
		ticks = append(ticks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get ticks: %w", err)
	}

	return ticks, nil
}

func NewDataStore(dsn string) (DataStore, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return DataStore{}, fmt.Errorf("error connecting to postgres: %w", err)
	}

	return DataStore{db: db}, nil
}
