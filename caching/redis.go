package caching

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	cacheTimeout = 50 * time.Millisecond
)

type RedisCachingService struct {
	client *redis.Client
}

func NewRedisCachingService(c *redis.Client) *RedisCachingService {
	return &RedisCachingService{
		client: c,
	}
}

func (svc *RedisCachingService) Get(ctx context.Context, key string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, cacheTimeout)
	defer cancel()

	val, err := svc.client.Get(ctx, key).Result()
	if err != nil {
		return "", nil
	}

	return val, nil
}

func (svc *RedisCachingService) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, cacheTimeout)
	defer cancel()

	_ = svc.client.Set(ctx, key, value, ttl)
	return nil
}

func (svc *RedisCachingService) Delete(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, cacheTimeout)
	defer cancel()

	_ = svc.client.Del(ctx, key)
	return nil
}
