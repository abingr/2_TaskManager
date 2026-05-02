package config

import "time"

const (
	RedisAddress      = "localhost:6379"
	TaskCacheTTL      = 5 * time.Minute
	RateLimitRequests = 5
	RateLimitWindow   = 10 * time.Second
)
