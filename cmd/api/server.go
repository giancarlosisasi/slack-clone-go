package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/giancarlosisasi/slack-clone-go/internal/config"
)

type Response struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

type Server struct {
	config *config.Config
	router *http.ServeMux
	server *http.Server
}

func NewServer(config *config.Config) *Server {
	router := http.NewServeMux()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Port),
		Handler:      router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTImeout,
		IdleTimeout:  config.IdleTimeout,
	}

	s := &Server{
		config: config,
		router: router,
		server: server,
	}

	s.registerRoutes()

	return s
}

func (s *Server) Start() error {
	log.Printf("Starting server on port %s\n", s.config.Port)

	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server...")
	return s.server.Shutdown(ctx)
}

func (s *Server) registerRoutes() {
	s.router.HandleFunc("/api/v1/health", s.withMiddleware(s.healthcheckHandler))
}

func (s *Server) withMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Log request
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Call the actual handler
		next(w, r)
	}
}

func (s *Server) sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (s *Server) sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := Response{
		Message: message,
		Status:  "error",
	}
	s.sendJSONResponse(w, response, statusCode)
}
