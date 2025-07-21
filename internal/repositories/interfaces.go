package repositories

type RoomStore interface {
	CreateRom(name string) (*Room, error)
	GetRoomByName(name string) (*Room, error)
	GetRoomByID(id string) (*Room, error)
	GetAllRooms() ([]*Room, error)
	AddUserToRoom(userID string, roomID string) error
	RemoveUserFromRoom(userID string, roomID string) error
}

type MessageStore interface {
	CreateMessage(userID string, roomID string, content string, messageType string) (*Message, error)
	GetRoomMessages(roomID string, limit int) ([]*Message, error)
	GetMessageWithUserInfo(roomID int, limit int) ([]*MessageWithUser, error)
	// GetMessagesSince(roomID int, since time.Time) ([]*Message, error)
}
