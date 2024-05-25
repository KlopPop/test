package sl

import (
	"log/slog"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "develop"
	EnvProd  = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case EnvDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	}
	return log
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
