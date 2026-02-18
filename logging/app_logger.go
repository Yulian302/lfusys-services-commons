package logger

import (
	"log/slog"
	"os"
	"strings"
)

func CreateAppLogger(env string) *slog.Logger {
	level := resolveLevel(env)

	opts := &slog.HandlerOptions{
		Level: level,
	}

	var handler slog.Handler

	if strings.EqualFold(env, "DEV") {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}

func resolveLevel(env string) slog.Level {
	if strings.EqualFold(env, "DEV") {
		return slog.LevelDebug
	}
	return slog.LevelInfo
}
