package retries

import (
	"context"
	"math/rand"
	"time"
)

const (
	DefaultAttempts  = 3
	DefaultBaseDelay = 100 * time.Millisecond
	HealthAttempts   = 3
	HealthBaseDelay  = 50 * time.Millisecond
)

// full jitter (AWS style): random in [0, exp)
func calculateDelayWithJitter(delay time.Duration, attempt int) time.Duration {
	exp := delay * time.Duration(1<<attempt)
	return time.Duration(rand.Int63n(int64(exp)))
}

func Retry(ctx context.Context, attempts int, baseDelay time.Duration, fn func() error, shouldRetry func(error) bool) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}

		if !shouldRetry(err) {
			return err
		}

		if i == attempts-1 {
			break
		}

		delay := calculateDelayWithJitter(baseDelay, i)

		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return err
}
