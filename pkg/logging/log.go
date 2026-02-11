package logging

import (
	"log/slog"
	"os"

	"github.com/m-bromo/my-game-list/config"
)

type Logger struct {
	Log *slog.Logger
}

func NewLogger(config *config.Config) *Logger {
	var l *slog.Logger

	switch config.Environment {
	case "staging":
		l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelWarn,
			AddSource: true,
		}))

	case "production":
		l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		}))

	default:
		l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: false,
		}))
	}

	return &Logger{
		Log: l,
	}
}
