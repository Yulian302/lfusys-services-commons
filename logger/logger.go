package logger

import (
	"log/slog"
	"os"
)

func CreateLogger(env string) *slog.Logger {
	var handler slog.Handler

	switch env {
	case "DEV":
		handler = PrettyHandler{
			h: slog.NewTextHandler(os.Stdout, nil),
		}
	default:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	return slog.New(handler)
}
