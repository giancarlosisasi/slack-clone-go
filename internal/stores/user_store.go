package stores

import (
	"context"

	database "github.com/giancarlosisasi/slack-clone-go/internal/database/generated"
	"github.com/giancarlosisasi/slack-clone-go/internal/models"
)

type UserPostgresStore struct {
	queries *database.Queries
}

func NewUserStore(queries *database.Queries) *UserPostgresStore {
	return &UserPostgresStore{
		queries: queries,
	}
}

func (s *UserPostgresStore) CreateUser(
	ctx context.Context,
	email string,
	firstName string,
	lastName string,
) (*models.User, error) {
	args := database.CreateUserParams{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
	user, err := s.queries.CreateUser(ctx, args)
	if err != nil {
		return nil, err
	}

	return s.mapCreateUserRowToModel(user), nil
}

func (s *UserPostgresStore) GetUserByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return s.mapGetUserByEmailRowToModel(user), err
}

func (s *UserPostgresStore) mapCreateUserRowToModel(dbUser database.CreateUserRow) *models.User {
	return &models.User{
		ID:        dbUser.ID.String(),
		Email:     dbUser.Email,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		UpdatedAt: dbUser.UpdatedAt.Time,
		CreatedAt: dbUser.CreatedAt.Time,
	}
}

func (s *UserPostgresStore) mapGetUserByEmailRowToModel(dbUser database.GetUserByEmailRow) *models.User {
	return &models.User{
		ID:        dbUser.ID.String(),
		Email:     dbUser.Email,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		UpdatedAt: dbUser.UpdatedAt.Time,
		CreatedAt: dbUser.CreatedAt.Time,
	}
}
