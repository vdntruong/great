package app

import (
	"context"
	"user-ms/internal/model"
)

// UserRepository aka User DAO
type UserRepository interface {
	Insert(ctx context.Context, email, username string, password string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, bool, error)
}
