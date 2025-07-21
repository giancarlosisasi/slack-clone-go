package main

import (
	"log"

	"github.com/giancarlosisasi/slack-clone-go/internal/app"
	"github.com/giancarlosisasi/slack-clone-go/internal/routes"
)

func main() {
	application, err := app.NewApplication()
	if err != nil {
		log.Fatalf("Failed to create application: %v", err)
	}

	defer application.DB.Close()

	routesHandler := routes.SetupRoutes(application)

	// Start the server with graceful shutdown
	// 1. Start http server
	// 2. Listen for shutdown signals (ctrl + c, Sigterm)
	// 3. Gracefully shutdown error
	// 4. Wait for background tasks to complete
	if err := application.Serve(routesHandler); err != nil {
		log.Fatalf("Server error: %v", err)
	}

	log.Println("Application shutdown successfully")
}
