package ratelimit

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRateLimiter struct {
	client *redis.Client
}

func NewRedisRateLimiter(c *redis.Client) *RedisRateLimiter {
	return &RedisRateLimiter{client: c}
}

func (r *RedisRateLimiter) Incr(ctx context.Context, key string) (int64, error) {
	res, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RedisRateLimiter) Expire(ctx context.Context, key string, ttl time.Duration) error {
	cmd := r.client.Expire(ctx, key, ttl)
	return cmd.Err()
}
