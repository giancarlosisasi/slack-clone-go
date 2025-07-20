package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Version      string
	Port         string
	AppEnv       string
	ReadTimeout  time.Duration
	WriteTImeout time.Duration
	IdleTimeout  time.Duration
}

func NewConfig() *Config {
	err := godotenv.Load(".env.backend")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:         os.Getenv("PORT"),
		Version:      os.Getenv("APP_VERSION"),
		AppEnv:       os.Getenv("APP_ENV"),
		ReadTimeout:  15 * time.Second,
		WriteTImeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
