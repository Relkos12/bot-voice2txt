package config

import (
	"errors"
	"log/slog"
	"os"
)

type Config struct {
	Token string
}

func LoadConfig(Logger *slog.Logger) (*Config, error) {

	const BOT_TOKEN_ENV = "BOT_TOKEN"

	botToken := os.Getenv(BOT_TOKEN_ENV)

	if botToken == "" {
		Logger.Error("Requred env variables are misssing", "BOT_TOKEN", botToken)
		return nil, errors.New("Requred env variables are missing")
	}

	return &Config{Token: botToken}, nil
}
