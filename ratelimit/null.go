package ratelimit

import (
	"context"
	"time"
)

type NullRateLimiter struct {
}

func NewNullRateLimiter() *NullRateLimiter {
	return &NullRateLimiter{}
}

func (r *NullRateLimiter) Incr(ctx context.Context, key string) (int64, error) {
	return 0, nil
}

func (r *NullRateLimiter) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return nil
}
