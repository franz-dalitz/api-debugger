package config

import (
	"log/slog"
	"os"
)

func parseSlogLevel(level string) slog.Leveler {
	levelmap := map[string]slog.Level{
		slog.LevelDebug.String(): slog.LevelDebug,
		slog.LevelInfo.String():  slog.LevelInfo,
		slog.LevelWarn.String():  slog.LevelWarn,
		slog.LevelError.String(): slog.LevelError,
	}
	return levelmap[level]
}

func ApplySlogSettings() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: parseSlogLevel(LogLevel),
	}))
	slog.SetDefault(logger)
	slog.Debug("applied settings to slog")
}
