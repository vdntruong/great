package services

import (
	"auth-ms/internal/entities/models"
	"context"
)

type IUserProvider interface {
	VerifyUser(ctx context.Context, email string, password string) (models.Credential, error)
}

type TokenAdaptor interface {
	GenerateTokenPair(userID string, username string) (string, string, error)
	ValidateToken(token string) error
}
