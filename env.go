package common

import "syscall"

func EnvVar(key string, fallback string) string {
	if val, ok := syscall.Getenv(key); ok && val != "" {
		return val
	}
	return fallback
}
