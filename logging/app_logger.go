package logger

import (
	"log/slog"
	"os"
	"strings"
)

func CreateAppLogger(env string) *slog.Logger {
	header := "\033[33mAPP \033[0m"
	level := resolveLevel(env)

	opts := &slog.HandlerOptions{
		Level: level,
	}

	if strings.EqualFold(env, "DEV") {
		return slog.New(NewPrettyHandler(os.Stdout, header))
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}

func resolveLevel(env string) slog.Level {
	if strings.EqualFold(env, "DEV") {
		return slog.LevelDebug
	}
	return slog.LevelInfo
}
