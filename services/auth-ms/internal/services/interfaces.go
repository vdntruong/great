package services

import (
	"context"

	"auth-ms/internal/entities/models"
)

type UserAdapter interface {
	VerifyUser(ctx context.Context, email string, password string) (models.Credential, error)
}

type TokenAdapter interface {
	GenerateTokenPair(userID string, username string) (string, string, error)
	ValidateToken(token string) error
}
