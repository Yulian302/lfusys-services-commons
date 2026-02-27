package config

import (
	"fmt"
	"strings"
)

type ServiceConfig struct {
	Gateway  GatewayConfig
	Sessions SessionsConfig
	Uploads  UploadsConfig
}

type GatewayConfig struct {
	Addr            string
	SessionsGRPCUrl string
	FrontendUrl     string
}

type SessionsConfig struct {
	Addr string
}

type UploadsConfig struct {
	Addr        string
	FrontendUrl string
}

func (c *ServiceConfig) Validate() error {
	var errs []string

	if err := c.Gateway.Validate(); err != nil {
		errs = append(errs, err.Error())
	}

	if err := c.Sessions.Validate(); err != nil {
		errs = append(errs, err.Error())
	}

	if err := c.Uploads.Validate(); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return fmt.Errorf("service config: %s", strings.Join(errs, "; "))
	}

	return nil
}

func (c *GatewayConfig) Validate() error {
	var errs []string

	if c.Addr == "" {
		errs = append(errs, "GATEWAY_ADDR is required")
	}

	if c.SessionsGRPCUrl == "" {
		errs = append(errs, "SESSION_GRPC_URL is required")
	}

	if c.FrontendUrl == "" {
		errs = append(errs, "FRONTEND_URL is required")
	}

	if len(errs) > 0 {
		return fmt.Errorf("gateway: %s", strings.Join(errs, "; "))
	}

	return nil
}

func (c *SessionsConfig) Validate() error {
	var errs []string

	if c.Addr == "" {
		errs = append(errs, "SESSIONS_GRPC_ADDR is required")
	}

	if len(errs) > 0 {
		return fmt.Errorf("sessions: %s", strings.Join(errs, "; "))
	}

	return nil
}

func (c *UploadsConfig) Validate() error {
	var errs []string

	if c.Addr == "" {
		errs = append(errs, "UPLOADS_ADDR is required")
	}

	if c.FrontendUrl == "" {
		errs = append(errs, "FRONTEND_URL is required")
	}

	if len(errs) > 0 {
		return fmt.Errorf("uploads: %s", strings.Join(errs, "; "))
	}

	return nil
}
