package retries

import (
	"context"
	"errors"

	"github.com/aws/smithy-go"
	"github.com/redis/go-redis/v9"
)

func IsContextError(err error) bool {
	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}

// this function checks if Redis error is transient (should be retried)
func IsRetriableRedisError(err error) bool {
	if err == nil {
		return false
	}

	if err == redis.Nil {
		return false
	}

	if IsContextError(err) {
		return false
	}

	return true
}

func IsRetriableDbError(err error) bool {
	if err == nil {
		return false
	}

	if IsContextError(err) {
		return false
	}

	var apiError smithy.APIError

	if errors.As(err, &apiError) {
		switch apiError.ErrorCode() {

		// throttling → retry
		case "ProvisionedThroughputExceededException",
			"ThrottlingException",
			"RequestLimitExceeded":
			return true

		// server faults → retry
		case "InternalServerError",
			"ServiceUnavailable":
			return true

		// client faults → no retry
		default:
			return false
		}

	}

	// network/transport errors (timeouts, EOF, etc)
	return true
}
