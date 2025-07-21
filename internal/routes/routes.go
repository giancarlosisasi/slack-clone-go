package routes

import (
	"net/http"

	"github.com/giancarlosisasi/slack-clone-go/internal/app"
	"github.com/giancarlosisasi/slack-clone-go/internal/handlers"
	"github.com/giancarlosisasi/slack-clone-go/internal/middleware"
)

func SetupRoutes(app *app.Application) http.Handler {
	router := http.NewServeMux()

	healthCheckHandler := handlers.NewHealthHandler(app)
	userHandler := handlers.NewUserHandler(app)

	router.HandleFunc("/api/v1/health", healthCheckHandler.HandleHealth)
	router.HandleFunc("/api/v1/user", userHandler.HandleRegisterUser)

	return middleware.Logging(router)
}
