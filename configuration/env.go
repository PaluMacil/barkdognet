package configuration

import (
	"log/slog"
	"strings"
)

type Env struct {
	Slug     string `koanf:"slug"`
	LogLevel string `koanf:"log_level"`
}

func (e Env) SlogLevel() slog.Level {
	var level slog.Level
	levelString := strings.ToLower(e.LogLevel)
	switch levelString {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "iarn":
		level = slog.LevelWarn
	case "Error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	return level
}
