package app

import (
	"log"
	"sync"

	logger "github.com/giancarlosisasi/slack-clone-go/internal"
	"github.com/giancarlosisasi/slack-clone-go/internal/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Logger *log.Logger
	Config *Config
	DB     *pgxpool.Pool
	WG     sync.WaitGroup
	// services
	// UserService *services.UserService
}

func NewApplication() (*Application, error) {
	config := NewConfig()
	logger := logger.NewLogger()

	db, err := database.New(logger)
	if err != nil {
		return nil, err
	}

	// stores
	// userStore := store.NewPostgresUserStore(db)

	// services
	// userService := services.NewUserService(userStore)

	return &Application{
		Logger: logger,
		Config: config,
		DB:     db,
		WG:     sync.WaitGroup{},
		// services
		// UserService: userService
	}, nil

}
