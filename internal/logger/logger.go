package logger

import (
	"log/slog"
	"os"
)

func Init() {
	var handler slog.Handler

	// Use JSON handler in production, Text handler in development
	// For now, let's use TextHandler for better readability in terminal
	handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
