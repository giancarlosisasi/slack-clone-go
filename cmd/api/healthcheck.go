package main

import (
	"net/http"
	"time"
)

func (s *Server) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Message: "Server is healthy and running!",
		Status:  "success",
		Data: map[string]interface{}{
			"timestamp": time.Now().UTC(),
			"version":   s.config.Version,
		},
	}

	s.sendJSONResponse(w, response, http.StatusOK)
}
