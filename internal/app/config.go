package app

import (
	"fmt"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	return &Config{
		Port:         fmt.Sprintf(":%s", port),
		Version:      os.Getenv("APP_VERSION"),
		AppEnv:       appEnv,
		ReadTimeout:  15 * time.Second,
		WriteTImeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
