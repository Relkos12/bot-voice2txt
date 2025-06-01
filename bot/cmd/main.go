package main

import (
	"log/slog"
	"os"
)

// setup logger
func setupLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	return logger
}

func main() {

}
