package repositories

import (
	"time"

	"github.com/giancarlosisasi/slack-clone-go/internal/models"
)

type Message struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	RoomID      string    `json:"room_id"`
	Content     string    `json:"content"`
	MessageType string    `json:"message_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MessageWithUser struct {
	Message Message     `json:"message"`
	User    models.User `json:"user"`
}
