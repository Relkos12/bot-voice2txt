package main

import (
	"log/slog"
	"os"

	"github.com/Relkos12/bot-voice2txt/internal/adapter/telegram"
	"github.com/Relkos12/bot-voice2txt/internal/config"
	"github.com/joho/godotenv"
)

// setup logger
func setupLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	return logger
}

func main() {

	if err := godotenv.Load(); err != nil {
		slog.Warn("env is not exist")
	}

	logger := setupLogger()

	//loading conf
	conf, err := config.LoadConfig(logger)
	if err != nil {
		logger.Error("error configuration not loaded")
		os.Exit(1)
	}

	//init adapter
	telegram.NewBotAdapter(conf.Token)

}
