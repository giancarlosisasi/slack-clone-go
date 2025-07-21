package handlers

import (
	"net/http"
	"time"

	"github.com/giancarlosisasi/slack-clone-go/internal/app"
)

type HealthcheckHandler struct {
	app *app.Application
}

func NewHealthHandler(app *app.Application) *HealthcheckHandler {
	return &HealthcheckHandler{
		app: app,
	}
}

func (h *HealthcheckHandler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		SendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Message: "Server is healthy and running!",
		Status:  "success",
		Data: map[string]interface{}{
			"timestamp": time.Now().UTC(),
			"version":   h.app.Config.Version,
		},
	}

	SendJSONResponse(w, response, http.StatusOK)
}
