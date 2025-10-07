package data

import "time"

type Tick struct {
	AssetID   string    `json:"assetID"`
	Timestamp time.Time `json:"timestamp"`
	LastPrice float64   `json:"lastPrice,omitempty"`
	LastSize  float64   `json:"lastSize,omitempty"`
	BestBid   float64   `json:"bestBid,omitempty"`
	BestAsk   float64   `json:"bestAsk,omitempty"`
}

/*
	// /ticks/{asset_id}?start={start}&end={end}&limit={limit}

id bigserial primary key
asset_id text
created_at timestamptz
last_price decimal
last_size decimal
best_bid decimal
best_ask decimal

index asset_id, created_at

type Candle struct {
	AssetID        string    `json:"assetID"`
	StartTimestamp time.Time `json:"timestamp"` start minute for which we calculate the candle
	Open           float64   `json:"open"` first
	High           float64   `json:"high"`
	Low            float64   `json:"low"`
	Close          float64   `json:"close"` last price of last tick
	Volume         float64   `json:"volume"` sum of the lastSize of each tick in the minute for asset_id
} //

// every minute first tick of the asset_id, last tick for the minute
*/
