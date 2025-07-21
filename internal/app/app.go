package app

import (
	"log"
	"sync"

	logger "github.com/giancarlosisasi/slack-clone-go/internal"
	"github.com/giancarlosisasi/slack-clone-go/internal/database"
	sqlcDatabase "github.com/giancarlosisasi/slack-clone-go/internal/database/generated"
	"github.com/giancarlosisasi/slack-clone-go/internal/repositories"
	"github.com/giancarlosisasi/slack-clone-go/internal/services"
	"github.com/giancarlosisasi/slack-clone-go/internal/stores"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Logger *log.Logger
	Config *Config
	DB     *pgxpool.Pool
	WG     sync.WaitGroup
	// services
	UserService *services.UserService
}

func NewApplication() (*Application, error) {
	config := NewConfig()
	logger := logger.NewLogger()

	db, err := database.New(logger)
	if err != nil {
		return nil, err
	}

	// sqlc queries
	queries := sqlcDatabase.New(db)

	// stores
	userStore := stores.NewUserStore(queries)

	// repository
	userRepository := repositories.NewUserRepository(userStore)

	// services
	userService := services.NewUserService(userRepository, logger)

	return &Application{
		Logger: logger,
		Config: config,
		DB:     db,
		WG:     sync.WaitGroup{},
		// services
		UserService: userService,
	}, nil

}
