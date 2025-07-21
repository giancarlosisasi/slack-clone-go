package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
}

func SendJSONResponse(w http.ResponseWriter, data any, statusCode int) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := Response{
		Message: message,
		Status:  "error",
	}
	SendJSONResponse(w, response, statusCode)
}
