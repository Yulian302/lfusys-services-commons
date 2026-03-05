package misc

import (
	"math/rand"
	"time"
)

func AddJitter(ttl time.Duration) time.Duration {
	jitterPercent := 0.1 // 10%
	maxJitter := time.Duration(float64(ttl) * jitterPercent)

	jitter := time.Duration(rand.Int63n(int64(maxJitter)))

	return ttl + jitter
}
