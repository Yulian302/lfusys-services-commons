package config

import (
	"fmt"
	"strings"
)

type SQSConfig struct {
	QueueName string
}

func (c *SQSConfig) Validate() error {
	var errs []string

	if c.QueueName == "" {
		errs = append(errs, "UPLOADS_NOTIFICATIONS_QUEUE_NAME is required")
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}

	return nil
}
