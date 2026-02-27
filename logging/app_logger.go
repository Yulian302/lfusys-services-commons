package logger

import (
	"log/slog"
	"os"

	"github.com/Yulian302/lfusys-services-commons/config"
)

func CreateAppLogger(env config.Environment) *slog.Logger {
	header := "\033[33mAPP \033[0m"
	level := resolveLevel(env)

	opts := &slog.HandlerOptions{
		Level: level,
	}

	if env == config.EnvDevelopment {
		return slog.New(NewPrettyHandler(os.Stdout, header))
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}

func resolveLevel(env config.Environment) slog.Level {
	if env == config.EnvDevelopment {
		return slog.LevelDebug
	}
	return slog.LevelInfo
}
