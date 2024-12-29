package app

import (
	"context"
	"user-ms/internal/model"
)

// UserRepository is an abstract interface to User of persistence mechanism, it's include the DAO and advantaged methods
type UserRepository interface {
	Insert(ctx context.Context, email, username string, password string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)

	FindByEmail(ctx context.Context, email string) (bool, error)
	FindByUsername(ctx context.Context, username string) (bool, error)
}
