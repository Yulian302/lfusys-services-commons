package logger

import (
	"log/slog"
	"os"
	"strings"
)

func CreateHttpLogger(env string) *slog.Logger {
	header := "\033[35mHTTP \033[0m"

	if strings.EqualFold(env, "DEV") {
		return slog.New(NewPrettyHandler(os.Stdout, header))
	}

	opts := &slog.HandlerOptions{Level: slog.LevelInfo}
	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}
