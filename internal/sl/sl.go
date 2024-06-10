package sl

import (
	"log/slog"
	"os"
)

const (
	envLocal   = "local"
	envDev     = "develop"
	envProd    = "prod"
	jsonFormat = "json"
	textFormat = "text"
)

func SetupLogger(env string, lvl int) *slog.Logger {
	var log *slog.Logger

	level := slog.Level(lvl)
	switch env {
	case textFormat:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	case jsonFormat:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	}
	return log
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
