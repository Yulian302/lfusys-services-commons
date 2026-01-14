package caching

import (
	"context"
	"time"
)

type NullCachingService struct {
	// do nothing
}

func NewNullCachingService() *NullCachingService {
	return &NullCachingService{}
}

func (svc *NullCachingService) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (svc *NullCachingService) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return nil
}

func (svc *NullCachingService) Delete(ctx context.Context, key string) error {
	return nil
}
