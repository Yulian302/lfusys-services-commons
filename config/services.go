package config

import "errors"

type ServiceConfig struct {
	UploadsAddr                   string
	UploadsNotificationsQueueName string
	GatewayAddr                   string
	SessionGRPCUrl                string
	SessionGRPCAddr               string
}

func (c *ServiceConfig) Validate() error {
	if c.UploadsAddr == "" {
		return errors.New("UPLOADS_ADDR is required")
	}

	if c.UploadsNotificationsQueueName == "" {
		return errors.New("UPLOADS_NOTIFICATION_QUEUE_NAME is required")
	}

	if c.GatewayAddr == "" {
		return errors.New("GATEWAY_ADDR is required")
	}

	if c.SessionGRPCUrl == "" {
		return errors.New("SESSION_GRPC_URL is required")
	}

	if c.SessionGRPCAddr == "" {
		return errors.New("SESSION_GRPC_ADDR is required")
	}

	return nil
}
