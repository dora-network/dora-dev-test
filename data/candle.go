package data

import "time"

type Candle struct {
	AssetID        string    `json:"assetID"`
	StartTimestamp time.Time `json:"timestamp"`
	Open           float64   `json:"open"`
	High           float64   `json:"high"`
	Low            float64   `json:"low"`
	Close          float64   `json:"close"`
	Volume         float64   `json:"volume"`
}

// /candles/{asset_id}?start={start}&end={end}&granularity={granularity}&limit={limit}
// every minute first tick of the asset_id, last tick for the minute
