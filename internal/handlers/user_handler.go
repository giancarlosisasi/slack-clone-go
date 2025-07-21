package handlers

import (
	"context"
	"net/http"

	"github.com/giancarlosisasi/slack-clone-go/internal/app"
)

type UserHandler struct {
	app *app.Application
}

func NewUserHandler(app *app.Application) *UserHandler {
	return &UserHandler{
		app: app,
	}
}

func (h *UserHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	// todo: make sure its post method request
	user, err := h.app.UserService.RegisterUser(
		context.Background(),
		"user-test@gmail.com",
		"user",
		"test",
	)

	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	SendJSONResponse(w, user, http.StatusCreated)
}
