package ratelimit

import (
	"context"
	"time"
)

type RateLimiter interface {
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, ttl time.Duration) error
}
