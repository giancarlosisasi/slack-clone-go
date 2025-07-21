package services

import (
	"context"
	"fmt"
	"log"

	"github.com/giancarlosisasi/slack-clone-go/internal/models"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, email string, firstName string, lastName string) (*models.User, error)
	// GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type UserService struct {
	repo   UserRepositoryInterface
	logger *log.Logger
}

func NewUserService(repo UserRepositoryInterface, logger *log.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, email string, firstName string, lastName string) (*models.User, error) {
	user, err := s.repo.CreateUser(ctx, email, firstName, lastName)
	if err != nil {
		s.logger.Printf("Failed to register user %s: %v\n", email, err)
		return nil, fmt.Errorf("registration failed: %w", err)
	}

	return user, nil
}
