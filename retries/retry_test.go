package retries

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRetry_SucceedsAfterFailures(t *testing.T) {
	attempts := 0

	err := Retry(
		context.Background(),
		3,
		1*time.Millisecond,
		func() error {
			attempts++
			if attempts < 3 {
				return errors.New("fail")
			}
			return nil
		},
		func(error) bool { return true },
	)

	require.NoError(t, err)
	require.Equal(t, 3, attempts)
}
