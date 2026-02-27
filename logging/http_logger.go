package logger

import (
	"log/slog"
	"os"

	"github.com/Yulian302/lfusys-services-commons/config"
)

func CreateHttpLogger(env config.Environment) *slog.Logger {
	header := "\033[35mHTTP \033[0m"

	if env == config.EnvDevelopment {
		return slog.New(NewPrettyHandler(os.Stdout, header))
	}

	opts := &slog.HandlerOptions{Level: slog.LevelInfo}
	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}
