package repositories

import (
	"context"
	"errors"

	"github.com/giancarlosisasi/slack-clone-go/internal/models"
)

type UserStoreInterface interface {
	CreateUser(ctx context.Context, email string, firstName string, lastName string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// GetUserById(id string) (*User, error)
	// GetRoomMembers(roomId string) ([]*User, error)
}

type UserRepository struct {
	store UserStoreInterface
}

func NewUserRepository(store UserStoreInterface) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, email string, firstName string, lastName string) (*models.User, error) {
	err := r.validateEmail(email)
	if err != nil {
		return nil, err
	}

	existingUser, err := r.store.GetUserByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, errors.New("username already exists")
	}

	return r.store.CreateUser(ctx, email, firstName, lastName)
}

func (r *UserRepository) validateEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	// todo: validate length, email format, etc

	return nil
}
