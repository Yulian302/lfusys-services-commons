package caching

import (
	"context"
	"time"

	logger "github.com/Yulian302/lfusys-services-commons/logging"
	"github.com/redis/go-redis/v9"
)

const (
	cacheTimeout = 50 * time.Millisecond
)

type RedisCachingService struct {
	client *redis.Client
	logger logger.Logger
}

func NewRedisCachingService(c *redis.Client, l logger.Logger) *RedisCachingService {
	return &RedisCachingService{
		client: c,
		logger: l,
	}
}

func (svc *RedisCachingService) Get(ctx context.Context, key string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, cacheTimeout)
	defer cancel()

	val, err := svc.client.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			svc.logger.Debug("cache get failed",
				"key", key,
				"error", err,
			)
		}
		return "", nil
	}

	return val, nil
}

func (svc *RedisCachingService) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, cacheTimeout)
	defer cancel()

	err := svc.client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		svc.logger.Debug("cache set failed",
			"key", key,
			"error", err,
		)
	}
	return err
}

func (svc *RedisCachingService) Delete(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, cacheTimeout)
	defer cancel()

	err := svc.client.Del(ctx, key).Err()
	if err != nil {
		svc.logger.Debug("cache delete failed",
			"key", key,
			"error", err,
		)
	}
	return err
}
