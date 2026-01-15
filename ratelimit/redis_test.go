package ratelimit

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestRedisRateLimiter(t *testing.T) {
	s, err := miniredis.Run()
	require.NoError(t, err)
	defer s.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	limiter := NewRedisRateLimiter(rdb)
	ctx := context.Background()
	key := "rate:test"

	// first increment should be 1
	count, err := limiter.Incr(ctx, key)
	require.NoError(t, err)
	require.Equal(t, int64(1), count)

	// should set ttl on first hit
	err = limiter.Expire(ctx, key, 1*time.Minute)
	require.NoError(t, err)

	ttl := s.TTL(key)
	require.True(t, ttl > 0)

	// next increment should be 2
	count, err = limiter.Incr(ctx, key)
	require.NoError(t, err)
	require.Equal(t, int64(2), count)
}
