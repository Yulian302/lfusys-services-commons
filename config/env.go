package config

import (
	"log"
	"strings"
	"syscall"
)

type Environment string

const (
	EnvProduction  Environment = "prod"
	EnvDevelopment Environment = "dev"
	EnvStaging     Environment = "staging"
	EnvTest        Environment = "test"
)

func EnvVar(key string, fallback string) string {
	if val, ok := syscall.Getenv(key); ok && val != "" {
		return val
	}
	return fallback
}

func ParseEnvironment(envRaw string) Environment {
	switch strings.ToLower(envRaw) {
	case "prod", "production":
		return EnvProduction
	case "dev", "development", "develop":
		return EnvDevelopment
	case "staging", "stage":
		return EnvStaging
	case "test", "testing":
		return EnvTest
	default:
		log.Printf("WARNING: Unknown environment '%s', defaulting to 'dev'", envRaw)
		return EnvDevelopment
	}
}
