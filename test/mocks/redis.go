package mocks

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// implements RedisClient interface
type FakeRedis struct {
	SetNXCalls int
	GetCalls   int
	DelCalls   int

	SetNXErrs []error
	GetErrs   []error
	DelErrs   []error

	Values map[string]string
}

func (f *FakeRedis) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	f.SetNXCalls++
	err := f.SetNXErrs[0]
	f.SetNXErrs = f.SetNXErrs[1:]
	cmd := redis.NewBoolResult(true, err)
	return cmd
}

func (f *FakeRedis) Get(ctx context.Context, key string) *redis.StringCmd {
	f.GetCalls++
	err := f.GetErrs[0]
	f.GetErrs = f.GetErrs[1:]
	cmd := redis.NewStringResult("true", err)
	return cmd
}

func (f *FakeRedis) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	f.DelCalls++
	err := f.DelErrs[0]
	f.DelErrs = f.DelErrs[1:]
	cmd := redis.NewIntResult(1, err)
	return cmd
}

func (f *FakeRedis) Close() error {
	return nil
}
