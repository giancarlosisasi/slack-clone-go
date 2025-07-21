package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(logger *log.Logger) (*pgxpool.Pool, error) {
	dbUrl := os.Getenv("DB_URL")
	pgxConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		logger.Printf("unable to parse the database url configuration: %v\n", err)
		return nil, err
	}

	pgxConfig.MaxConns = 2
	pgxConfig.MaxConnIdleTime = time.Minute * 15

	if dbUrl == "" {
		panic("DB_URL env variable can't be empty")
	}
	dbpool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		logger.Printf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	err = dbpool.Ping(context.Background())
	if err != nil {
		logger.Printf("Ping to database failed: %v\n", err)
		os.Exit(1)
	}

	return dbpool, nil
}
