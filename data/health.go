package data

import "time"

type Health struct {
	Uptime time.Duration `json:"uptime"`
	// map of endpoint to number of requests
	Requests map[string]int64 `json:"requests"`
	// map of endpoint to number of failed requests
	Failures map[string]int64 `json:"failures"`
}
